package main

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"net/http"
)

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.log.
		Err(err).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("Bad request")

	err = writeJSONError(w, http.StatusBadRequest, err.Error())
	if err != nil {
		app.log.
			Err(err).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("Error while json error")
	}
}

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.log.
		Err(err).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("Internal server error")

	err = writeJSONError(w, http.StatusInternalServerError, err.Error())
	if err != nil {
		app.log.
			Err(err).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("Error while json error")
	}
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.log.
		Err(err).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("Unauthorized error")

	err = writeJSONError(w, http.StatusUnauthorized, err.Error())
	if err != nil {
		app.log.
			Err(err).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("Error while json error")
	}
}

func (app *application) notFoundErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.log.
		Err(err).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("Not found error")

	err = writeJSONError(w, http.StatusNotFound, err.Error())
	if err != nil {
		app.log.
			Err(err).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("Error while json error")
	}
}

func (app *application) forbiddenErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.log.
		Err(err).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("Forbidden error")

	err = writeJSONError(w, http.StatusForbidden, err.Error())
	if err != nil {
		app.log.
			Err(err).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg("Error while json error")
	}
}

func (app *application) checkRole(w http.ResponseWriter, r *http.Request, role enum.UserRole) bool {
	user := getUserFromRequest(r)

	ok := user.Role.Level() >= role.Level()

	if !ok {
		app.forbiddenErrorResponse(w, r, errors.New("you do not have permissions"))
	}

	return ok
}
