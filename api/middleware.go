package api

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zagvozdeen/examly/internal/store"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"time"
)

func (app *Application) loggerMiddleware(next http.Handler) http.Handler {
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

func (app *Application) authMiddleware(next http.Handler) http.Handler {
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
			//app.internalServerError(w, r, fmt.Errorf("parse token: %w, header: %s, parts: %v", err, header, parts))
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
			if errors.Is(err, store.ErrNotFound) {
				app.notFoundErrorResponse(w, r, err)
				return
			}
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
