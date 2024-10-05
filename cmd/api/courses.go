package main

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *application) getCourses(w http.ResponseWriter, r *http.Request) {
	var filter store.GetCoursesFilter
	query := r.URL.Query()
	if query.Has("user_id") {
		id, err := strconv.Atoi(query.Get("user_id"))
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		filter.UserID = id
	}
	if query.Has("trashed") {
		trashed, err := strconv.ParseBool(query.Get("trashed"))
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		filter.Trashed = trashed
	}
	if query.Has("statuses") {
		statuses := query["statuses"]
		filter.Statuses = make([]any, len(statuses))
		for i, status := range statuses {
			if _, err := enum.NewStatus(status); err != nil {
				app.badRequestResponse(w, r, errors.New("invalid status"))
				return
			}
			filter.Statuses[i] = status
		}
	}

	courses, err := app.store.CoursesStore.Get(r.Context(), filter)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": courses,
	})
}

type CreateCoursePayload struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=2048"`
	Color       string `json:"color" validate:"required,max=255"`
	Icon        string `json:"icon" validate:"required,max=255"`
}

func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {
	var payload CreateCoursePayload
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

	course := &store.Course{
		UUID:        uid.String(),
		Name:        payload.Name,
		Description: payload.Description,
		Color:       payload.Color,
		Icon:        payload.Icon,
		Status:      enum.CreatedStatus,
	}

	ctx := r.Context()

	if err := app.store.CoursesStore.Create(ctx, course); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, course)
}

func (app *application) getCourse(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()

	course, err := app.store.CoursesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, err)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	questions, err := app.store.QuestionsStore.GetByCourseID(ctx, course.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"course":    course,
		"questions": questions,
	})
}

//func (app *application) name(w http.ResponseWriter, r *http.Request) {}
