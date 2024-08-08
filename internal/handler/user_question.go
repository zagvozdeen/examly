package handler

import (
	"encoding/json"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) checkAnswer(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	if uuid == "" {
		http.Error(w, "empty uuid", http.StatusBadRequest)
		return
	}

	var input service.CheckAnswerInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.QuestionUUID = uuid

	question, err := h.services.UserQuestions.CheckAnswer(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": question,
	})
}
