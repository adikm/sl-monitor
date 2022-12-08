package notifications

import (
	"net/http"
	request "sl-monitor/internal/server"
	"sl-monitor/internal/server/auth"
)

func Routes(nh *Handler) {
	http.HandleFunc("/notifications", request.MustBe(http.MethodPost, auth.MustBeLoggedIn(nh.Create)))
}
