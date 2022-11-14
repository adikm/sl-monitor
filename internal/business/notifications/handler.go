package notifications

import (
	"fmt"
	"net/http"
	"sl-monitor/internal"
	request "sl-monitor/internal/server"
	"time"
)

type NotificationHandler struct {
	Notifications *NotificationStore
	common        *internal.JsonCommon
}

func NewHandler(store *NotificationStore, common *internal.JsonCommon) *NotificationHandler {
	return &NotificationHandler{store, common}
}

func (nh *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		nh.common.MethodNotAllowed(w, r)
		return
	}

	var input struct {
		Email     string               `json:"email"`
		Timestamp time.Time            `json:"timestamp"`
		Weekdays  internal.WeekdaysSum `json:"weekdays"`
	}

	err := request.DecodeJSON(r, &input)
	if err != nil {
		nh.common.BadRequest(w, r, err)
		return
	}

	newNot, err := nh.Notifications.Create(input.Email, input.Timestamp, input.Weekdays)

	fmt.Println(newNot)

	if err != nil {
		nh.common.ServerError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
