package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) getCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.services.GetCourses()
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
