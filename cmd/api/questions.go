package main

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/guregu/null/v5"
	"net/http"
	"time"
)

func (app *application) getQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := app.store.QuestionsStore.Get(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": questions,
	})
}

type CreateQuestionPayload struct {
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:""`
	Explanation string `json:"explanation" validate:""`
	Type        string `json:"type" validate:"required"`
	CourseID    int    `json:"course_id" validate:"required"`
	FileID      int    `json:"file_id" validate:""`
	ModuleID    int    `json:"module_id" validate:""`
	Answers     []struct {
		ID        int    `json:"id" validate:"required"`
		Content   string `json:"content" validate:"required"`
		IsCorrect bool   `json:"is_correct" validate:""`
	} `json:"answers" validate:"required,dive,required"`
}

func (app *application) createQuestion(w http.ResponseWriter, r *http.Request) {
	var payload CreateQuestionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	uid, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	t, err := enum.NewQuestionType(payload.Type)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := getUserFromRequest(r)

	var options store.Options
	for _, answer := range payload.Answers {
		options = append(options, store.Option{
			ID:        answer.ID,
			Content:   answer.Content,
			IsCorrect: answer.IsCorrect,
		})
	}

	question := &store.Question{
		UUID:        uid.String(),
		Title:       payload.Title,
		Content:     null.StringFrom(payload.Content),
		Explanation: null.StringFrom(payload.Explanation),
		Type:        t,
		Status:      enum.CreatedStatus,
		CourseID:    payload.CourseID,
		ModuleID:    null.IntFrom(int64(payload.ModuleID)),
		CreatedBy:   user.ID,
		Options:     options,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := app.store.QuestionsStore.Create(r.Context(), question); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": question,
	})
}

type UpdateQuestionPayload struct {
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:""`
	Explanation string `json:"explanation" validate:""`
	Type        string `json:"type" validate:"required"`
	CourseID    int    `json:"course_id" validate:"required"`
	FileID      int    `json:"file_id" validate:""`
	ModuleID    int    `json:"module_id" validate:""`
	Answers     []struct {
		ID        int    `json:"id" validate:"required"`
		Content   string `json:"content" validate:"required"`
		IsCorrect bool   `json:"is_correct" validate:""`
	} `json:"answers" validate:"required,dive,required"`
}

func (app *application) updateQuestion(w http.ResponseWriter, r *http.Request) {
	var payload UpdateQuestionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	question, err := app.store.QuestionsStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if question.CreatedBy != user.ID {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to update this question"))
		return
	}

	newUUID, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	t, err := enum.NewQuestionType(payload.Type)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var options store.Options
	for _, answer := range payload.Answers {
		options = append(options, store.Option{
			ID:        answer.ID,
			Content:   answer.Content,
			IsCorrect: answer.IsCorrect,
		})
	}

	nextQuestion := &store.Question{
		UUID:           newUUID.String(),
		Title:          payload.Title,
		Content:        null.StringFrom(payload.Content),
		Explanation:    null.StringFrom(payload.Explanation),
		Type:           t,
		Status:         enum.CreatedStatus,
		CourseID:       payload.CourseID,
		ModuleID:       null.IntFrom(int64(payload.ModuleID)),
		PrevQuestionID: null.IntFrom(int64(question.ID)),
		CreatedBy:      user.ID,
		Options:        options,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = app.store.QuestionsStore.Create(ctx, nextQuestion)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	question.NextQuestionID = null.IntFrom(int64(nextQuestion.ID))
	question.UpdatedAt = time.Now()

	err = app.store.QuestionsStore.Update(ctx, &question)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": question,
	})
}
