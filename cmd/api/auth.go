package main

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type RegisterUserPayload struct {
	FirstName            string `json:"first_name" validate:"required,max=255"`
	LastName             string `json:"last_name" validate:"required,max=255"`
	Email                string `json:"email" validate:"required,email,max=255"`
	Password             string `json:"password" validate:"required,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	uid, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	user := &store.User{
		UUID:      uid.String(),
		Email:     null.StringFrom(payload.Email),
		FirstName: null.StringFrom(payload.FirstName),
		LastName:  null.StringFrom(payload.LastName),
		Role:      enum.MemberRole,
		Password:  null.StringFrom(string(bytes)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.store.UsersStore.Create(r.Context(), user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": user,
	})
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.UsersStore.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			app.unauthorizedErrorResponse(w, r, errors.New("invalid email or password"))
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(payload.Password))
	if err != nil {
		app.unauthorizedErrorResponse(w, r, errors.New("invalid email or password"))
		return
	}

	token, err := generateTokenByUserID(app.config.SecretKey, user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]string{
		"data": token,
	})
}

func (app *application) getGuestToken(w http.ResponseWriter, r *http.Request) {
	uid, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	user := &store.User{
		UUID:      uid.String(),
		Role:      enum.GuestRole,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.store.UsersStore.Create(r.Context(), user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	token, err := generateTokenByUserID(app.config.SecretKey, user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]string{
		"data": token,
	})
}

func generateTokenByUserID(secret string, id int) (string, error) {
	claims := &claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
