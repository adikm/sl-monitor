package stations

import (
	"net/http"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/server/response"
)

type Handler struct {
	service trafikverket.Service
}

func NewHandler(service trafikverket.Service) *Handler {
	return &Handler{service}
}

func (sh *Handler) FetchStations(w http.ResponseWriter, r *http.Request) {
	stations, err := sh.service.FetchStations()

	if err != nil {
		response.ServerError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, stations)
	if err != nil {
		response.ServerError(w, r, err)
	}
}
