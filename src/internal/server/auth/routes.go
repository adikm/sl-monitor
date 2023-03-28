package auth

import (
	"net/http"
	"sl-monitor/internal/server"
)

func Routes(ah *Handler) {

	http.HandleFunc("/login", request.MustBe(http.MethodPost, ah.login))
	http.HandleFunc("/logout", request.MustBe(http.MethodPost, MustBeLoggedIn(ah.Logout)))

}
