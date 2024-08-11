package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

func (h *Handler) getMyCourses(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	courses, err := h.services.GetCoursesByUserID(user.ID)
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

func (h *Handler) getAllCourses(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	courses, err := h.services.GetAllCourses(user.ID)
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

func (h *Handler) createCourse(w http.ResponseWriter, r *http.Request) {
	var u service.CreateCourseInput
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

	user := r.Context().Value("user").(*model.User)

	id, err := h.services.Courses.CreateCourse(user, &u)
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

func (h *Handler) getCourseByUUID(w http.ResponseWriter, r *http.Request) {
	uuid, found := mux.Vars(r)["uuid"]
	if !found {
		encode(w, r, http.StatusBadRequest, fmt.Errorf("empty uuid"))
		return
	}

	course, err := h.services.Courses.GetCourseByUUID(uuid)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	user := current(r)
	stats, err := h.services.Courses.GetStatsByUserID(user.ID)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	errorCourse, err := h.services.UserCourses.GetUserCourseByTypeAndUserID(model.ErrorUserCourseType, user.ID)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, CustomResponse{
		"data":   course,
		"stats":  stats,
		"errors": errorCourse,
	})
}

func (h *Handler) getCourseStats(w http.ResponseWriter, r *http.Request) {
	uuid, found := mux.Vars(r)["uuid"]
	if !found {
		encode(w, r, http.StatusBadRequest, fmt.Errorf("empty uuid"))
		return
	}

	course, err := h.services.Courses.GetCourseByUUID(uuid)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	user := current(r)
	stats, err := h.services.Courses.GetCourseStatsByUUID(&model.CourseStatsParams{
		CourseID: course.ID,
		UserID:   user.ID,
	})
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, stats)
}

func (h *Handler) getStats(w http.ResponseWriter, r *http.Request) {
	user := current(r)
	stats, err := h.services.Courses.GetCourseStatsByUUID(&model.CourseStatsParams{
		UserID: user.ID,
	})
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, stats)
}

func (h *Handler) createMarathon(w http.ResponseWriter, r *http.Request) {
	user := current(r)

	uuid := mux.Vars(r)["uuid"]
	if uuid == "" {
		encode(w, r, http.StatusBadRequest, fmt.Errorf("empty uuid"))
		return
	}

	input := &service.CreateUserCourseInput{
		CourseUUID: uuid,
		UserID:     user.ID,
		Type:       model.MarathonUserCourseType,
	}

	uuid, err := h.services.Courses.CreateUserCourse(input)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusCreated, uuid)
}

func (h *Handler) createExam(w http.ResponseWriter, r *http.Request) {
	user := current(r)

	uuid := mux.Vars(r)["uuid"]
	if uuid == "" {
		encode(w, r, http.StatusBadRequest, fmt.Errorf("empty uuid"))
		return
	}

	input := &service.CreateUserCourseInput{
		CourseUUID: uuid,
		UserID:     user.ID,
		Type:       model.ExamUserCourseType,
	}

	uuid, err := h.services.Courses.CreateUserCourse(input)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusCreated, uuid)
}
