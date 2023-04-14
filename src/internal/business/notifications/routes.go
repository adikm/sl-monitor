package notifications

import (
	"github.com/go-chi/chi/v5"
	"sl-monitor/internal/server/auth"
)

func Routes(r *chi.Mux, nh *Handler) {
	r.Post("/notifications", auth.MustBeLoggedIn(nh.create))
	r.Get("/notifications/all", auth.MustBeLoggedIn(nh.findForCurrentUser))
}
