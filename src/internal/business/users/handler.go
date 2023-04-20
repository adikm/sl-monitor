package users

import (
	"net/http"
	"sl-monitor/internal/server"
	"sl-monitor/internal/server/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (uh *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input UserRequest

	err := request.DecodeJSON(r, &input)
	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	exists, err := uh.service.UserExists(input.Email)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}

	if exists {
		response.ErrorMessage(w, http.StatusBadRequest, "User with given e-mail already exists", nil)
		return
	}

	u, err := uh.service.Create(input)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, u)
	if err != nil {
		response.ServerError(w, r, err)
	}
}
