package users

import (
	"net/http"
	"sl-monitor/internal/server"
)

func Routes(uh *Handler) {
	http.HandleFunc("/users", request.MustBe(http.MethodPost, uh.create))
}
