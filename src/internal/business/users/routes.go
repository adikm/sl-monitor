package users

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, uh *Handler) {
	r.Post("/users", uh.create)
}
