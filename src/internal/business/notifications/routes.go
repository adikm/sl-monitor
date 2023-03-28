package notifications

import (
	"net/http"
	"sl-monitor/internal/server"
	"sl-monitor/internal/server/auth"
)

func Routes(nh *Handler) {
	http.HandleFunc("/notifications", request.MustBe(http.MethodPost, auth.MustBeLoggedIn(nh.create)))
	http.HandleFunc("/notifications/all", request.MustBe(http.MethodGet, auth.MustBeLoggedIn(nh.findForCurrentUser)))
}
