package handler

import (
	"encoding/json"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/go-playground/validator/v10"
	"net/http"
	"slices"
)

func (h *Handler) getQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.services.GetQuestions()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": questions,
	})
}

func (h *Handler) getMyQuestions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	questions, err := h.services.GetQuestionsByUserID(user.ID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": questions,
	})
}

func (h *Handler) createQuestion(w http.ResponseWriter, r *http.Request) {
	var u service.CreateQuestionInput
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

	if !slices.Contains(model.AllQuestionTypes, u.Type) {
		http.Error(w, "invalid question type", http.StatusBadRequest)
		return
	}

	if len(u.Answers) == 0 {
		http.Error(w, "at least one answer is required", http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(*model.User)

	id, err := h.services.Questions.CreateQuestion(user, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": id,
	})
}
