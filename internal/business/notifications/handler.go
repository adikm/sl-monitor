package notifications

import (
	"net/http"
	"sl-monitor/internal"
	request "sl-monitor/internal/server"
	"sl-monitor/internal/server/response"
	"time"
)

type Handler struct {
	Notifications Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store}
}

func (nh *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.MethodNotAllowed(w, r)
		return
	}

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

	id, err := nh.Notifications.Create(input.Email, input.Timestamp, input.Weekdays)
	if err != nil {
		response.ServerError(w, r, err)
		return
	}

	n := &Notification{id, input.Email, input.Timestamp, input.Weekdays.AsWeekdays()}

	err = response.JSON(w, http.StatusOK, n)
	if err != nil {
		response.ServerError(w, r, err)
	}
}
