package stations

import (
	"net/http"
	"sl-monitor/internal"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server/response"
)

type StationHandler struct {
	config *config.Config
	common *internal.JsonCommon
}

func NewHandler(config *config.Config, common *internal.JsonCommon) *StationHandler {
	return &StationHandler{config, common}
}

func (sh *StationHandler) FetchStations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sh.common.MethodNotAllowed(w, r)
		return
	}

	stations, err := trafikverket.FetchStations(sh.config.TrafficAPI.AuthKey)

	if err != nil {
		sh.common.ServerError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, stations)
	if err != nil {
		sh.common.ServerError(w, r, err)
	}
}
