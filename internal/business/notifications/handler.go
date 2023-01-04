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
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store}
}

func (nh *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email     string               `json:"email"`
		Timestamp time.Time            `json:"timestamp"`
		Weekdays  internal.WeekdaysSum `json:"weekdays"`
	}

	err := request.DecodeJSON(r, &input)
	if err != nil {
		response.BadRequest(w, r, err)
		return
	}
	userId := getUserId(r)
	id, err := nh.store.Create(input.Timestamp, input.Weekdays, userId)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}

	n := &Notification{id, input.Timestamp, input.Weekdays.AsWeekdays(), userId}

	err = response.JSON(w, http.StatusOK, n)
	if err != nil {
		response.ServerError(w, r, err)
	}
}

func (nh *Handler) FindAllForWeekday(weekday internal.Weekday) ([]Notification, error) {
	return nh.store.FindAll(weekday)
}

func (nh *Handler) findForCurrentUser(w http.ResponseWriter, r *http.Request) {
	userId := getUserId(r)
	notifications, err := nh.store.FindByUserId(userId)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, notifications)
	if err != nil {
		response.ServerError(w, r, err)
	}
}

func getUserId(r *http.Request) int {
	cookie, _ := r.Cookie("session_token")
	sessionToken := cookie.Value
	return auth.Sessions[sessionToken].UserId
}
