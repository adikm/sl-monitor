package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, ah *Handler) {

	r.Post("/login", ah.login)
	r.Post("/logout", MustBeLoggedIn(ah.Logout))

}
