package main

import (
	"net/http"
	"sl-monitor/internal"
	request "sl-monitor/internal/server"
	"sl-monitor/internal/server/response"
	"sl-monitor/internal/trafikverket"
	"time"
)

func (app *application) stations(w http.ResponseWriter, r *http.Request) {
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

func (app *application) createNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		app.methodNotAllowed(w, r)
		return
	}

	var input struct {
		Email     string               `json:"email"`
		Timestamp time.Time            `json:"timestamp"`
		Weekdays  internal.WeekdaysSum `json:"weekdays"`
	}

	err := request.DecodeJSON(r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	err = app.db.CreateNotification(input.Email, input.Timestamp, input.Weekdays)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
