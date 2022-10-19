package main

import (
	"sl-monitor/client"
	"sl-monitor/config"
)

func main() {
	config.Load()

	client.FetchDepartures()
}
