package handler

import (
	"encoding/json"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) getModules(w http.ResponseWriter, r *http.Request) {
	courses, err := h.services.GetModules()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": courses,
	})
}

func (h *Handler) createModule(w http.ResponseWriter, r *http.Request) {
	var u service.CreateModuleInput
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.Modules.CreateModule(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": id,
	})
}
