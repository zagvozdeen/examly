package main

import (
	"context"
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type application struct {
	log    zerolog.Logger
	config config
	store  store.Storage
}

type config struct {
	AppEnv    string
	AppURL    string
	DBAddr    string
	SecretKey string
}

func (app *application) mount() *mux.Router {
	router := mux.NewRouter()

	router.Use(app.loggerMiddleware)

	router.HandleFunc("/", app.viewIndex)

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/auth/register", app.register).Methods("POST")
	subRouter.HandleFunc("/auth/login", app.login).Methods("POST")
	subRouter.HandleFunc("/auth/guest-token", app.getGuestToken).Methods("GET")

	authRouter := subRouter.PathPrefix("/").Subrouter()
	authRouter.Use(app.authMiddleware)

	authRouter.HandleFunc("/me", app.getCurrentUser).Methods("GET")
	authRouter.HandleFunc("/me", app.roleMiddleware(enum.MemberRole, app.getCurrentUser)).Methods("PATCH")

	authRouter.HandleFunc("/courses", app.getCourses).Methods("GET")
	authRouter.HandleFunc("/courses", app.roleMiddleware(enum.MemberRole, app.createCourse)).Methods("POST")
	authRouter.HandleFunc("/courses/{uuid}", app.getCourse).Methods("GET")

	authRouter.HandleFunc("/modules", app.getModules).Methods("GET")
	authRouter.HandleFunc("/modules", app.roleMiddleware(enum.MemberRole, app.createModule)).Methods("POST")

	authRouter.HandleFunc("/questions", app.getQuestions).Methods("GET")
	authRouter.HandleFunc("/questions", app.roleMiddleware(enum.MemberRole, app.createQuestion)).Methods("POST")
	authRouter.HandleFunc("/questions/{uuid}", app.roleMiddleware(enum.MemberRole, app.updateQuestion)).Methods("PATCH")

	authRouter.HandleFunc("/files", app.uploadFile).Methods("POST")

	authRouter.HandleFunc("/test-sessions", app.createTestSession).Methods("POST")
	authRouter.HandleFunc("/test-sessions/stats", app.getUserStats).Methods("GET")
	authRouter.HandleFunc("/test-sessions/{uuid}", app.getTestSession).Methods("GET")

	authRouter.HandleFunc("/user-answers", app.checkAnswer).Methods("POST")

	authRouter.HandleFunc("/users", app.getUsers).Methods("GET")

	return router
}

func (app *application) run(router *mux.Router) error {
	server := &http.Server{
		Addr:         app.config.AppURL,
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		app.log.Info().Str("signal", s.String()).Msg("Signal caught")

		shutdown <- server.Shutdown(ctx)
	}()

	app.log.Info().Str("addr", app.config.AppURL).Msg("Server has started")

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.log.Info().Str("addr", app.config.AppURL).Msg("Server has stopped")

	return nil
}
