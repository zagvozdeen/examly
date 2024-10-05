package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestUserRegister(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()

	t.Run("check response if body is empty", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("check response if passwords is different", func(t *testing.T) {
		p := RegisterUserPayload{
			FirstName:            "tet",
			LastName:             "tes",
			Email:                "fsdfs@fsdf.fsdf",
			Password:             "123",
			PasswordConfirmation: "1234",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("check response if body not fulfilled", func(t *testing.T) {
		p := RegisterUserPayload{
			FirstName: "tet",
			LastName:  "tes",
			Password:  "123",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("check response if all data is correct", func(t *testing.T) {
		p := RegisterUserPayload{
			FirstName:            "tet",
			LastName:             "tes",
			Email:                "fsdfs@fsdf.fsdf",
			Password:             "password",
			PasswordConfirmation: "password",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusCreated, rr.Code)
	})

	t.Run("check response if email is not unique", func(t *testing.T) {
		p := RegisterUserPayload{
			FirstName:            "tet",
			LastName:             "tes",
			Email:                "guest@mail.ru",
			Password:             "password",
			PasswordConfirmation: "password",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestGetGuestToken(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()

	t.Run("check response code is 200", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodGet, "/api/v1/auth/guest-token", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}

func TestLogin(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()

	t.Run("check response code if body is empty", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("check response is ok", func(t *testing.T) {
		p := LoginPayload{
			Email:    "guest@mail.ru",
			Password: "password",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
	})

	t.Run("check response if password is incorrect", func(t *testing.T) {
		p := LoginPayload{
			Email:    "guest@mail.ru",
			Password: "sword",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("check response if user does not exists", func(t *testing.T) {
		p := LoginPayload{
			Email:    "notguest@mail.ru",
			Password: "password",
		}
		body, _ := json.Marshal(p)

		r, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(r, mux)

		checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	})
}
