package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type CustomResponse map[string]any

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var response map[string]any
	switch val := any(v).(type) {
	case error:
		response = map[string]any{
			"message": val.Error(),
		}
	case CustomResponse:
		response = val
	default:
		response = map[string]any{
			"data": v,
		}
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error(err.Error())
	}
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		return v, err
	}
	return v, nil
}

func current(r *http.Request) *model.User {
	return r.Context().Value("user").(*model.User)
}
