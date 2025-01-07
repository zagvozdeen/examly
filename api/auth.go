package api

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/zagvozdeen/examly/internal/enum"
	"github.com/zagvozdeen/examly/internal/store"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type RegisterUserPayload struct {
	Role                 string `json:"role" validate:"required"`
	FirstName            string `json:"first_name" validate:"required,max=255"`
	LastName             string `json:"last_name" validate:"required,max=255"`
	Email                string `json:"email" validate:"required,email,max=255"`
	Password             string `json:"password" validate:"required,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	role, err := enum.NewUserRole(payload.Role)
	if err != nil || !(role == enum.MemberRole || role == enum.ReferralRole || role == enum.CompanyRole) {
		app.badRequestResponse(w, r, errors.New("invalid role"))
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

	ctx := r.Context()
	_, err = app.store.UsersStore.GetByEmail(ctx, payload.Email)
	if !errors.Is(err, store.ErrNotFound) {
		app.badRequestResponse(w, r, errors.New("email already exists"))
		return
	}

	user := &store.User{
		UUID:      uid.String(),
		Email:     null.StringFrom(payload.Email),
		FirstName: null.StringFrom(payload.FirstName),
		LastName:  null.StringFrom(payload.LastName),
		Role:      role,
		Password:  null.StringFrom(string(bytes)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.store.UsersStore.Create(ctx, user)
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

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
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
			app.badRequestResponse(w, r, errors.New("invalid email or password"))
		} else {
			app.internalServerError(w, r, err)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(payload.Password))
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid email or password"))
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

func (app *Application) getGuestToken(w http.ResponseWriter, r *http.Request) {
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
