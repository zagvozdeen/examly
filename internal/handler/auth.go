package handler

import (
	"github.com/Den4ik117/examly/internal/service"
	"net/http"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	u, err := decode[service.RegisterInput](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Auth.CreateUser(u)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusCreated, id)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	u, err := decode[service.LoginInput](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, err)
		return
	}

	jwt, err := h.services.Auth.Login(u)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, jwt)
}

func (h *Handler) getGuestToken(w http.ResponseWriter, r *http.Request) {
	u, err := h.services.GetGuestToken()
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, u)
}

func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) {
	user := current(r)

	encode(w, r, http.StatusOK, user)
}

func (h *Handler) updateMe(w http.ResponseWriter, r *http.Request) {
	u, err := decode[service.UpdateUserInput](r)
	user := current(r)

	updated, err := h.services.Auth.UpdateUser(*user, &u)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, updated)
}

func (h *MuxHandler) register(w http.ResponseWriter, r *http.Request) {
	u, err := decode[service.RegisterInput](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Auth.CreateUser(u)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusCreated, id)
}

func (h *MuxHandler) login(w http.ResponseWriter, r *http.Request) {
	u, err := decode[service.LoginInput](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, err)
		return
	}

	jwt, err := h.services.Auth.Login(u)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, jwt)
}

func (h *MuxHandler) getGuestToken(w http.ResponseWriter, r *http.Request) {
	u, err := h.services.GetGuestToken()
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, u)
}

func (h *MuxHandler) getMe(w http.ResponseWriter, r *http.Request) {
	user := current(r)

	encode(w, r, http.StatusOK, user)
}

func (h *MuxHandler) updateMe(w http.ResponseWriter, r *http.Request) {
	u, err := decode[service.UpdateUserInput](r)
	user := current(r)

	updated, err := h.services.Auth.UpdateUser(*user, &u)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, updated)
}
