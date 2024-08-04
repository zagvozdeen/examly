package handler

import (
	"context"
	"net/http"
)

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		if len([]rune(t)) <= 7 {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}
		t = t[7:]

		user, err := h.services.Auth.CheckAuth(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) guestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		if len([]rune(t)) <= 8 {
			next.ServeHTTP(w, r)
			return
		}
		t = t[8:]

		_, err := h.services.Auth.CheckAuth(t)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "user is not a guest", http.StatusForbidden)
	})
}
