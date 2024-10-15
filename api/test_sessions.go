package api

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/guregu/null/v5"
	"math/rand"
	"net/http"
	"time"
)

type CreateTestSessionPayload struct {
	CourseUUID string `json:"course_uuid" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Shuffle    bool   `json:"shuffle" validate:""`
}

func (app *Application) createTestSession(w http.ResponseWriter, r *http.Request) {
	var payload CreateTestSessionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
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

	course, err := app.store.CoursesStore.GetByUUID(ctx, payload.CourseUUID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	questions, err := app.store.QuestionsStore.GetByCourseID(ctx, course.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var ids []int
	for _, question := range questions {
		ids = append(ids, question.ID)
	}
	if payload.Shuffle {
		rand.Shuffle(len(ids), func(i, j int) {
			ids[i], ids[j] = ids[j], ids[i]
		})
	}

	user := getUserFromRequest(r)

	test := &store.TestSession{
		UUID:        uid.String(),
		Name:        course.Name,
		Type:        t,
		UserID:      user.ID,
		CourseID:    null.IntFrom(int64(course.ID)),
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

func (app *Application) getUserStats(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(r)

	stats, err := app.store.TestSessionsStore.GetStats(r.Context(), user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": stats,
	})
}

func (app *Application) getTestSession(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("uuid is required"))
		return
	}

	user := getUserFromRequest(r)

	test, err := app.store.TestSessionsStore.GetByUUID(r.Context(), uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if test.UserID != user.ID {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to access this test session"))
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": test,
	})
}

//func (app *Application) name(w http.ResponseWriter, r *http.Request) {}
