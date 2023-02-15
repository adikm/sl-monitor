package notifications

import (
	"net/http"
	"sl-monitor/internal"
	request "sl-monitor/internal/server"
	"sl-monitor/internal/server/auth"
	"sl-monitor/internal/server/response"
	"time"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (nh *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email       string               `json:"email"`
		Timestamp   time.Time            `json:"timestamp"`
		Weekdays    internal.WeekdaysSum `json:"weekdays"`
		StationCode string               `json:"stationCode"`
	}

	err := request.DecodeJSON(r, &input)
	if err != nil {
		response.BadRequest(w, r, err)
		return
	}
	userId := currentUserId(r)
	n, err := nh.service.Create(input.Timestamp, input.Weekdays, userId, input.StationCode)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, n)
	if err != nil {
		response.ServerError(w, r, err)
	}
}

func (nh *Handler) findForCurrentUser(w http.ResponseWriter, r *http.Request) {
	userId := currentUserId(r)
	notifications, err := nh.service.findByUserId(userId)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, notifications)
	if err != nil {
		response.ServerError(w, r, err)
	}
}

func currentUserId(r *http.Request) int {
	cookie, _ := r.Cookie("session_token")
	sessionToken := cookie.Value
	return auth.Sessions[sessionToken].UserId
}
