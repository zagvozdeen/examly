package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/zagvozdeen/examly/internal/store"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	log    zerolog.Logger
	config Config
	store  store.Storage
}

type Config struct {
	IsProduction bool
	AppURL       string
	DBAddr       string
	SecretKey    string
}

func NewApplication(log zerolog.Logger, config Config, store store.Storage) *Application {
	return &Application{
		log:    log,
		config: config,
		store:  store,
	}
}

func (app *Application) Mount() *mux.Router {
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
	authRouter.HandleFunc("/me", app.updateCurrentUser).Methods("PATCH")

	authRouter.HandleFunc("/courses", app.getCourses).Methods("GET")
	authRouter.HandleFunc("/courses", app.createCourse).Methods("POST")
	authRouter.HandleFunc("/courses/{uuid}", app.getCourse).Methods("GET")
	authRouter.HandleFunc("/courses/{uuid}", app.updateCourse).Methods("PATCH")
	authRouter.HandleFunc("/courses/{uuid}", app.deleteCourse).Methods("DELETE")
	authRouter.HandleFunc("/courses/{uuid}/moderate", app.moderateCourse).Methods("PATCH")

	authRouter.HandleFunc("/modules", app.getModules).Methods("GET")
	authRouter.HandleFunc("/modules", app.createModule).Methods("POST")
	authRouter.HandleFunc("/modules/{uuid}", app.getModule).Methods("GET")
	authRouter.HandleFunc("/modules/{uuid}", app.updateModule).Methods("PATCH")
	authRouter.HandleFunc("/modules/{uuid}", app.deleteModule).Methods("DELETE")
	authRouter.HandleFunc("/modules/{uuid}/moderate", app.moderateModule).Methods("PATCH")

	authRouter.HandleFunc("/questions", app.getQuestions).Methods("GET")
	authRouter.HandleFunc("/questions", app.createQuestion).Methods("POST")
	authRouter.HandleFunc("/questions/{uuid}", app.getQuestion).Methods("GET")
	authRouter.HandleFunc("/questions/{uuid}", app.updateQuestion).Methods("PATCH")
	authRouter.HandleFunc("/questions/{uuid}", app.deleteQuestion).Methods("DELETE")
	authRouter.HandleFunc("/questions/{uuid}/moderate", app.moderateQuestion).Methods("PATCH")

	authRouter.HandleFunc("/files", app.uploadFile).Methods("POST")

	authRouter.HandleFunc("/test-sessions", app.getTestSessions).Methods("GET")
	authRouter.HandleFunc("/test-sessions", app.createTestSession).Methods("POST")
	authRouter.HandleFunc("/test-sessions/{uuid}", app.getTestSession).Methods("GET")

	authRouter.HandleFunc("/user-answers", app.checkAnswer).Methods("POST")

	authRouter.HandleFunc("/users", app.getUsers).Methods("GET")
	authRouter.HandleFunc("/users/experience", app.getUserExperience).Methods("GET")
	authRouter.HandleFunc("/users/experience", app.createUserExperience).Methods("POST")

	authRouter.HandleFunc("/tags", app.getTags).Methods("GET")

	return router
}

func (app *Application) Run(router *mux.Router) error {
	server := &http.Server{
		Addr:         app.config.AppURL,
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	initValidator()

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

	app.log.Info().Str("addr", fmt.Sprintf("http://%s", app.config.AppURL)).Msg("Server has started")

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
