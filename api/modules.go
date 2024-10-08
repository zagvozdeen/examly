package api

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

func (app *Application) getModules(w http.ResponseWriter, r *http.Request) {
	modules, err := app.store.ModulesStore.Get(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": modules,
	})
}

type CreateModulePayload struct {
	Name     string `json:"name" validate:"required,max=255"`
	CourseID int    `json:"course_id" validate:"required"`
}

func (app *Application) createModule(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	var payload CreateModulePayload
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

	user := getUserFromRequest(r)

	module := &store.Module{
		UUID:      uid.String(),
		Name:      payload.Name,
		Status:    enum.CreatedStatus,
		CourseID:  payload.CourseID,
		CreatedBy: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := app.store.ModulesStore.Create(r.Context(), module); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": module,
	})
}

func (app *Application) getModule(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("missing uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	module, err := app.store.ModulesStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if module.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to view this module"))
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": module,
	})
}

type UpdateModulePayload struct {
	Name string `json:"name" validate:"required,max=255"`
}

func (app *Application) updateModule(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("missing uuid"))
		return
	}

	var payload UpdateModulePayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := getUserFromRequest(r)
	ctx := r.Context()

	module, err := app.store.ModulesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, errors.New("module not found"))
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}
	if module.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to update this module"))
		return
	}

	module.Name = payload.Name
	module.UpdatedAt = time.Now()

	err = app.store.ModulesStore.Update(r.Context(), &module)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": module,
	})
}

func (app *Application) deleteModule(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("missing uuid"))
		return
	}

	user := getUserFromRequest(r)
	ctx := r.Context()

	module, err := app.store.ModulesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, errors.New("module not found"))
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}
	if module.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to delete this module"))
		return
	}

	module.DeletedAt = null.TimeFrom(time.Now())
	module.UpdatedAt = time.Now()

	err = app.store.ModulesStore.Delete(r.Context(), &module)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type ModerateModulePayload struct {
	Status           string `json:"status" validate:"required"`
	ModerationReason string `json:"moderation_reason" validate:""`
}

func (app *Application) moderateModule(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.ModeratorRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("missing uuid"))
		return
	}

	var payload ModerateModulePayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := getUserFromRequest(r)
	ctx := r.Context()

	module, err := app.store.ModulesStore.GetByUUID(ctx, uid)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.notFoundErrorResponse(w, r, errors.New("module not found"))
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

	module.ModerationReason = null.StringFrom(payload.ModerationReason)
	module.ModeratedBy = null.IntFrom(int64(user.ID))
	module.UpdatedAt = time.Now()
	module.Status = s

	err = app.store.ModulesStore.UpdateStatus(r.Context(), &module)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": module,
	})
}
