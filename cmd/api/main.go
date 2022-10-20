package main

import (
	"sl-monitor/internal/client"
	"sl-monitor/internal/config"
)

func main() {
	config.Load()

	client.FetchDepartures()
	client.FetchStations()
}
