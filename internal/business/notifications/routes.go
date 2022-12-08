package notifications

import (
	"net/http"
	"sl-monitor/internal/server/auth"
)

func Routes(nh *Handler) {
	http.HandleFunc("/notifications", auth.MustBeLoggedIn(nh.Create))
}
