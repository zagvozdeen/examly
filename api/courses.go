package api

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/guregu/null/v5"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) getCourses(w http.ResponseWriter, r *http.Request) {
	var filter store.GetCoursesFilter
	query := r.URL.Query()
	user := getUserFromRequest(r)
	ctx := r.Context()

	if query.Has("created_by") {
		id, err := strconv.Atoi(query.Get("created_by"))
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		if id != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
			app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
			return
		}
		filter.CreatedBy = id
	}
	if query.Has("or_created_by") {
		id, err := strconv.Atoi(query.Get("or_created_by"))
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		if id != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
			app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
			return
		}
		filter.OrCreatedBy = id
	}
	if query.Has("all") {
		if user.Role.Level() < enum.ModeratorRole.Level() {
			app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
			return
		}
		filter.All = true
	}

	courses, err := app.store.CoursesStore.Get(ctx, filter)
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

func (app *Application) createCourse(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

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

	ctx := r.Context()
	user := getUserFromRequest(r)

	course := &store.Course{
		UUID:        uid.String(),
		Name:        payload.Name,
		Description: payload.Description,
		Color:       payload.Color,
		Icon:        payload.Icon,
		Status:      enum.CreatedStatus,
		CreatedBy:   user.ID,
	}

	if err := app.store.CoursesStore.Create(ctx, course); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": course,
	})
}

func (app *Application) getCourse(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	course, err := app.store.CoursesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, err)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}
	if course.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
		return
	}

	questions, err := app.store.QuestionsStore.GetByCourseID(ctx, course.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data":      course,
		"questions": questions,
	})
}

type UpdateCoursePayload struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=2048"`
	Color       string `json:"color" validate:"required,max=255"`
	Icon        string `json:"icon" validate:"required,max=255"`
}

func (app *Application) updateCourse(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	var payload UpdateCoursePayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	course, err := app.store.CoursesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, err)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}
	if course.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
		return
	}

	course.Name = payload.Name
	course.Description = payload.Description
	course.Color = payload.Color
	course.Icon = payload.Icon
	course.UpdatedAt = time.Now()

	err = app.store.CoursesStore.Update(ctx, &course)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": course,
	})
}

func (app *Application) deleteCourse(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	course, err := app.store.CoursesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, err)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}
	if course.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
		return
	}

	err = app.store.CoursesStore.Delete(ctx, &course)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type ModerateCoursePayload struct {
	ModerationReason string `json:"moderation_reason" validate:"max=1024"`
	Status           string `json:"status" validate:"required"`
}

func (app *Application) moderateCourse(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.ModeratorRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	var payload ModerateCoursePayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	course, err := app.store.CoursesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, err)
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	s, err := enum.NewStatus(payload.Status)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	course.Status = s
	course.ModeratedBy = null.IntFrom(int64(user.ID))
	course.ModerationReason = null.NewString(payload.ModerationReason, payload.ModerationReason != "")
	course.UpdatedAt = time.Now()

	err = app.store.CoursesStore.UpdateStatus(ctx, &course)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": course,
	})
}
