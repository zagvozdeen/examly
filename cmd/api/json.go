package main

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var Validate *validator.Validate

func initValidator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	return writeJSON(w, status, map[string]string{
		"error": message,
	})
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	if r.Body == nil {
		return errors.New("body is empty")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func (app *application) jsonResponse(w http.ResponseWriter, r *http.Request, status int, data any) {
	err := writeJSON(w, status, data)
	if err != nil {
		app.log.
			Err(err).
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Msg("Error writing JSON response")
	}
}
