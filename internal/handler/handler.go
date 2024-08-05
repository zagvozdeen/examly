package handler

import (
	"github.com/Den4ik117/examly/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", h.viewIndex).Methods("GET")

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	guestRouter := subRouter.PathPrefix("/").Subrouter()
	guestRouter.Use(h.guestMiddleware)

	authRouter := subRouter.PathPrefix("/").Subrouter()
	authRouter.Use(h.authMiddleware)

	guestRouter.HandleFunc("/auth/register", h.register).Methods("POST")
	guestRouter.HandleFunc("/auth/login", h.login).Methods("POST")
	guestRouter.HandleFunc("/auth/guest-token", h.getGuestToken).Methods("GET")

	authRouter.HandleFunc("/me", h.getMe).Methods("GET")
	authRouter.HandleFunc("/me", h.updateMe).Methods("PATCH")

	authRouter.HandleFunc("/courses", h.getCourses).Methods("GET")
	authRouter.HandleFunc("/courses", h.createCourse).Methods("POST")
	authRouter.HandleFunc("/courses/{uuid}", h.getCourseByUUID).Methods("GET")

	authRouter.HandleFunc("/modules", h.getModules).Methods("GET")
	authRouter.HandleFunc("/modules", h.createModule).Methods("POST")

	authRouter.HandleFunc("/questions", h.getQuestions).Methods("GET")
	authRouter.HandleFunc("/questions", h.createQuestion).Methods("POST")

	authRouter.HandleFunc("/files", h.uploadFile).Methods("POST")

	return router
}
