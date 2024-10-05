package main

import (
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (app *application) getModules(w http.ResponseWriter, r *http.Request) {
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

func (app *application) createModule(w http.ResponseWriter, r *http.Request) {
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
