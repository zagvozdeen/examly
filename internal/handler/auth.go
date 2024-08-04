package handler

import (
	"encoding/json"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var u service.RegisterInput
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

	id, err := h.services.Auth.CreateUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": id,
	})
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var u service.LoginInput
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

	jwt, err := h.services.Auth.Login(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": jwt,
	})
}

func (h *Handler) getGuestToken(w http.ResponseWriter, r *http.Request) {
	u, err := h.services.GetGuestToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": u,
	})
}

func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": user,
	})
}
