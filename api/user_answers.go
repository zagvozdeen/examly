package api

import (
	"encoding/json"
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"net/http"
	"slices"
	"time"
)

type CheckAnswerPayload struct {
	TestSessionID int    `json:"course_id" validate:"required"`
	QuestionID    int    `json:"question_id" validate:"required"`
	AnswerID      int    `json:"answer_id" validate:""`
	AnswersIDs    []int  `json:"answers_ids" validate:""`
	Plaintext     string `json:"plaintext" validate:""`
}

func (app *Application) checkAnswer(w http.ResponseWriter, r *http.Request) {
	var payload CheckAnswerPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	data, err := json.Marshal(map[string]any{
		"answer_id":   payload.AnswerID,
		"answers_ids": payload.AnswersIDs,
		"plaintext":   payload.Plaintext,
	})
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	test, err := app.store.TestSessionsStore.GetByID(ctx, payload.TestSessionID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, errors.New("test session not found"))
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	if test.UserID != user.ID {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to answer this question"))
		return
	}

	if !slices.Contains(test.QuestionIDs, payload.QuestionID) {
		app.notFoundErrorResponse(w, r, errors.New("question not found in the test session"))
		return
	}

	question, err := app.store.QuestionsStore.GetByID(ctx, payload.QuestionID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, errors.New("question not found"))
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	if question.Type == enum.PlaintextQuestionType && payload.Plaintext == "" {
		app.badRequestResponse(w, r, errors.New("plaintext question must have plaintext answer"))
		return
	}

	if question.Type == enum.SingleChoiceQuestionType && payload.AnswerID == 0 {
		app.badRequestResponse(w, r, errors.New("single choice question must have answer id"))
		return
	}

	if question.Type == enum.MultipleChoiceQuestionType && payload.AnswersIDs == nil {
		app.badRequestResponse(w, r, errors.New("multiple choice question must have answers ids"))
		return
	}

	var correct bool

	if payload.AnswerID != 0 {
		for _, option := range question.Options {
			correct = option.IsCorrect && option.ID == payload.AnswerID
		}
	}

	if payload.AnswersIDs != nil {
		correct = true
		for _, option := range question.Options {
			for _, id := range payload.AnswersIDs {
				if option.ID == id && !option.IsCorrect {
					correct = false
					break
				}
			}
		}
	}

	if payload.Plaintext != "" {
		for _, option := range question.Options {
			correct = option.IsCorrect && option.Content == payload.Plaintext
		}
	}

	answer := &store.UserAnswer{
		TestSessionID: payload.TestSessionID,
		QuestionID:    payload.QuestionID,
		AnswerData:    string(data),
		IsCorrect:     correct,
		AnsweredAt:    time.Now(),
	}

	err = app.store.UserAnswersStore.Create(ctx, answer)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if !correct {
		mistake, err := app.store.TestSessionsStore.GetByUserIDAndType(ctx, user.ID, enum.MistakeTestSessionType)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				uid, err := uuid.NewV7()
				if err != nil {
					app.internalServerError(w, r, err)
					return
				}

				mistake = store.TestSession{
					UUID:        uid.String(),
					Name:        "Mistakes",
					Type:        enum.MistakeTestSessionType,
					UserID:      user.ID,
					QuestionIDs: []int{question.ID},
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				err = app.store.TestSessionsStore.Create(ctx, &mistake)
				if err != nil {
					app.internalServerError(w, r, err)
					return
				}
			} else {
				app.internalServerError(w, r, err)
				return
			}
		} else {
			mistake.QuestionIDs = append(mistake.QuestionIDs, question.ID)
			err := app.store.TestSessionsStore.Update(ctx, &mistake)
			if err != nil {
				app.internalServerError(w, r, err)
				return
			}
		}
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": answer,
	})
}

//func (app *Application) name(w http.ResponseWriter, r *http.Request) {}
