package api

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/guregu/null/v5"
	"github.com/zagvozdeen/examly/internal/enum"
	"github.com/zagvozdeen/examly/internal/store"
	"math/rand"
	"net/http"
	"time"
)

func (app *Application) getTestSessions(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(r)
	filter := store.GetTestSessionsFilter{UserID: user.ID}
	query := r.URL.Query()
	ctx := r.Context()

	if query.Has("course_uuid") {
		course, err := app.store.CoursesStore.GetByUUID(ctx, query.Get("course_uuid"))
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}
		filter.CourseID = course.ID
	}

	sessions, err := app.store.TestSessionsStore.Get(ctx, filter)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ids := make([]int, len(sessions))
	for i, session := range sessions {
		ids[i] = session.ID
	}
	stats, err := app.store.TestSessionsStore.GetStats(ctx, ids)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	for _, stat := range stats {
		for i, session := range sessions {
			if session.ID == stat.ID {
				sessions[i].Correct = stat.Correct
				sessions[i].Incorrect = stat.Incorrect
			}
		}
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": sessions,
	})
}

type CreateTestSessionPayload struct {
	CourseUUID string `json:"course_uuid"`
	Type       string `json:"type"`
	Shuffle    bool   `json:"shuffle"`
	TagsIDs    []int  `json:"tags_ids"`
}

func (app *Application) createTestSession(w http.ResponseWriter, r *http.Request) {
	var payload CreateTestSessionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	t, err := enum.NewTestSessionType(payload.Type)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if t == enum.MistakeTestSessionType {
		app.badRequestResponse(w, r, errors.New("mistake test session is not allowed"))
		return
	}

	uid, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	var ids []int
	var name string
	var courseId null.Int
	if t == enum.SelectionSystemTestSessionType {
		if len(payload.TagsIDs) == 0 {
			app.badRequestResponse(w, r, errors.New("tags are required for selection system test session"))
			return
		}
		ids, err = app.store.TagsStore.GetTagsQuestions(ctx, payload.TagsIDs)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}
		name = "Тест системы подбора"
	} else {
		course, err := app.store.CoursesStore.GetByUUID(ctx, payload.CourseUUID)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				app.notFoundErrorResponse(w, r, err)
				return
			}
			app.internalServerError(w, r, err)
			return
		}

		questions, err := app.store.QuestionsStore.GetByCourseID(ctx, course.ID)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		for _, question := range questions {
			ids = append(ids, question.ID)
		}
		name = course.Name
		courseId = null.IntFrom(int64(course.ID))
	}

	if payload.Shuffle {
		rand.Shuffle(len(ids), func(i, j int) {
			ids[i], ids[j] = ids[j], ids[i]
		})
	}

	user := getUserFromRequest(r)

	test := &store.TestSession{
		UUID:        uid.String(),
		Name:        name,
		Type:        t,
		UserID:      user.ID,
		CourseID:    courseId,
		QuestionIDs: ids,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.store.TestSessionsStore.Create(ctx, test)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": test,
	})
}

func (app *Application) getTestSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("uuid is required"))
		return
	}

	user := getUserFromRequest(r)

	test, err := app.store.TestSessionsStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	test.Questions, err = app.store.QuestionsStore.GetByIDs(ctx, test.QuestionIDs)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	answers, err := app.store.UserAnswersStore.GetByTestSessionID(ctx, test.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	for _, answer := range answers {
		for i, question := range test.Questions {
			if answer.QuestionID == question.ID {
				test.Questions[i].UserAnswers = append(test.Questions[i].UserAnswers, answer)
				break
			}
		}
	}

	if test.UserID != user.ID {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to access this test session"))
		return
	}

	ua, err := app.store.UserAnswersStore.GetByTestSessionID(ctx, test.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	uam := make(map[int]store.UserAnswer, len(ua))
	for _, answer := range ua {
		uam[answer.QuestionID] = answer
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data":    test,
		"answers": uam,
	})
}
