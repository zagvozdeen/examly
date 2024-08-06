package handler

import (
	"encoding/json"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/service"
	"github.com/Den4ik117/examly/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/guregu/null/v5"
	"log"
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

	user := r.Context().Value("user").(*model.User)
	id, err := h.services.Modules.CreateModule(user, &u)
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

func (h *Handler) getMyModules(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	modules, err := h.services.GetModulesByUserID(user.ID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	courses, err := h.services.GetModuleCourses(modules)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	m := util.GetMapByIDFromSlice(courses)
	for i, module := range modules {
		modules[i].Course = null.ValueFrom(m[module.CourseID])
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": modules,
	})
}

func (h *Handler) getAllModules(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)
	modules, err := h.services.GetAllModules(user.ID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data": modules,
	})
}
