package stations

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, sh *Handler) {
	r.Get("/stations", sh.FetchStations)
}
