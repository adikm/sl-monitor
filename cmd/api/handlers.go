package main

import (
	"net/http"
	"sl-monitor/internal/server/response"
	"sl-monitor/internal/trafikverket"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		app.methodNotAllowed(w, r)
		return
	}

	stations, err := trafikverket.FetchStations(app.config.TrafficAPI.AuthKey)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, stations)
	if err != nil {
		app.serverError(w, r, err)
	}
}
