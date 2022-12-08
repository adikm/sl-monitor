package stations

import (
	"net/http"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server/response"
)

type Handler struct {
	config  *config.Config
	service trafikverket.Service
}

func NewHandler(config *config.Config, service trafikverket.Service) *Handler {
	return &Handler{config, service}
}

func (sh *Handler) FetchStations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		response.MethodNotAllowed(w, r)
		return
	}

	stations, err := sh.service.FetchStations(sh.config.TrafficAPI.AuthKey)

	if err != nil {
		response.ServerError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, stations)
	if err != nil {
		response.ServerError(w, r, err)
	}
}
