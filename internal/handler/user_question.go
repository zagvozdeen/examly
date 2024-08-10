package handler

import (
	"fmt"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) checkAnswer(w http.ResponseWriter, r *http.Request) {
	uuid, found := mux.Vars(r)["uuid"]
	if !found {
		encode(w, r, http.StatusBadRequest, fmt.Errorf("empty uuid"))
		return
	}

	input, err := decode[service.CheckAnswerInput](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, err)
		return
	}
	user := current(r)
	input.QuestionUUID = uuid
	input.UserID = user.ID

	question, err := h.services.UserQuestions.CheckAnswer(&input)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, question)
}
