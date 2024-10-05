package main

import (
	"github.com/den4ik117/examly/internal/store"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	initValidator()

	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "02.01.2006 15:04:05",
	}).With().Timestamp().Logger()

	storage := store.NewMockStorage(t)

	return &application{
		log: logger,
		config: config{
			AppEnv:    "testing",
			AppURL:    "127.0.0.1:8888",
			SecretKey: "qwert12345",
		},
		store: storage,
	}
}

func executeRequest(r *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, r)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
