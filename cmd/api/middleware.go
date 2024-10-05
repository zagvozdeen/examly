package main

import (
	"errors"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"time"
)

func (app *application) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		latency := time.Since(start)
		app.log.
			Debug().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("latency", latency.String()).
			Msg("Request received")
	})
}

func (app *application) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			app.unauthorizedErrorResponse(w, r, errors.New("missing authorization header"))
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, errors.New("incorrect authorization header"))
			return
		}

		claims := &claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (any, error) {
			return []byte(app.config.SecretKey), nil
		})
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}
		if !token.Valid {
			app.unauthorizedErrorResponse(w, r, errors.New("token does not valid"))
			return
		}

		ctx := r.Context()

		user, err := app.store.UsersStore.GetByID(ctx, claims.UserID)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromRequest(r *http.Request) store.User {
	return r.Context().Value("user").(store.User)
}

func (app *application) roleMiddleware(role enum.UserRole, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)

		if user.Role.Level() < role.Level() {
			app.forbiddenErrorResponse(w, r, errors.New("you do not have permissions"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (app *application) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)

		if user.Role == enum.AdminRole || user.Role == enum.ModeratorRole {
			next.ServeHTTP(w, r)
			return
		}

		app.forbiddenErrorResponse(w, r, errors.New("you do not have permissions"))
	})
}
