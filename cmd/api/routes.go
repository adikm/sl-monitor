package main

import (
	"net/http"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/stations"
)

func (app *application) routes(nh *notifications.NotificationHandler, sh *stations.StationHandler) {

	http.HandleFunc("/", app.jsonCommon.NotFound)
	http.HandleFunc("/stations", sh.FetchStations)
	http.HandleFunc("/notifications", nh.CreateNotification)

}
