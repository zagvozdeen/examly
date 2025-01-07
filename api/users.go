package api

import (
	"errors"
	"fmt"
	"github.com/guregu/null/v5"
	"github.com/zagvozdeen/examly/internal/enum"
	"net/http"
)

func (app *Application) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(r)

	if user.Role.Level() > enum.GuestRole.Level() {
		user.FullName = null.StringFrom(fmt.Sprintf("%s %s", user.LastName.String, user.FirstName.String))
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": user,
	})
}

type UpdateUserPayload struct {
	Role        string `json:"role" validate:"required"`
	FirstName   string `json:"first_name" validate:"required,max=255"`
	LastName    string `json:"last_name" validate:"required,max=255"`
	Email       string `json:"email" validate:"required,email,max=255"`
	Description string `json:"description" validate:""`
	CompanyName string `json:"company_name" validate:""`
	Contact     string `json:"contact" validate:""`
}

func (app *Application) updateCurrentUser(w http.ResponseWriter, r *http.Request) {
	var payload UpdateUserPayload
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

	role, err := enum.NewUserRole(payload.Role)
	if err != nil || !(role == enum.MemberRole || role == enum.ReferralRole || role == enum.CompanyRole) {
		app.badRequestResponse(w, r, errors.New("invalid role"))
		return
	}
	if user.Role.Level() >= enum.ModeratorRole.Level() {
		role = user.Role
	}

	user.Role = role
	user.FirstName = null.StringFrom(payload.FirstName)
	user.LastName = null.StringFrom(payload.LastName)
	user.Email = null.StringFrom(payload.Email)
	user.FullName = null.StringFrom(fmt.Sprintf("%s %s", user.LastName.String, user.FirstName.String))
	user.Description = null.NewString(payload.Description, payload.Description != "")
	user.CompanyName = null.NewString(payload.CompanyName, payload.CompanyName != "")
	user.Contact = null.NewString(payload.Contact, payload.Contact != "")

	err = app.store.UsersStore.Update(ctx, &user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": user,
	})
}

func (app *Application) getUsers(w http.ResponseWriter, r *http.Request) {
	u := getUserFromRequest(r)
	if u.Role.Level() < enum.ModeratorRole.Level() && u.Role != enum.CompanyRole {
		app.forbiddenErrorResponse(w, r, errors.New("you do not have permissions"))
		return
	}

	users, err := app.store.UsersStore.Get(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	for i, user := range users {
		if user.LastName.IsZero() && user.FirstName.IsZero() {
			user.FullName = null.StringFrom(fmt.Sprintf("Гость #%d", user.ID))
		} else {
			user.FullName = null.StringFrom(fmt.Sprintf("%s %s", user.LastName.String, user.FirstName.String))
		}
		users[i] = user
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": users,
	})
}
