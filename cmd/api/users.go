package main

import (
	"github.com/den4ik117/examly/internal/enum"
	"github.com/guregu/null/v5"
	"net/http"
)

func (app *application) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(r)

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": user,
	})
}

type UpdateUserPayload struct {
	FirstName string `json:"first_name" validate:"required,max=255"`
	LastName  string `json:"last_name" validate:"required,max=255"`
	Email     string `json:"email" validate:"required,email,max=255"`
}

func (app *application) updateCurrentUser(w http.ResponseWriter, r *http.Request) {
	var payload UpdateUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := getUserFromRequest(r)
	user.FirstName = null.StringFrom(payload.FirstName)
	user.LastName = null.StringFrom(payload.LastName)
	user.Email = null.StringFrom(payload.Email)

	err := app.store.UsersStore.Update(r.Context(), &user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": user,
	})
}

func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.ModeratorRole); !ok {
		return
	}

	users, err := app.store.UsersStore.Get(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": users,
	})
}