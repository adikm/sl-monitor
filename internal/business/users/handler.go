package users

import (
	"net/http"
	request "sl-monitor/internal/server"
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
	n, err := uh.service.Create(input)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, n)
	if err != nil {
		response.ServerError(w, r, err)
	}
}
