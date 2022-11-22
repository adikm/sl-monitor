package notifications

import (
	"net/http"
	"sl-monitor/internal/server/auth"
)

func Routes(nh *NotificationHandler) {
	http.HandleFunc("/notifications", auth.MustBeLoggedIn(nh.CreateNotification))
}
