package auth

import "net/http"

func Routes(ah *AuthHandler) {

	http.HandleFunc("/login", ah.login)
	http.HandleFunc("/logout", MustBeLoggedIn(ah.Logout))

}
