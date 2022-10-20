package main

import (
	"fmt"
	"sl-monitor/internal/client"
	"sl-monitor/internal/config"
)

func main() {
	config.Load()

	departues := client.FetchDepartures()
	stations := client.FetchStations()
	fmt.Println(stations)
	fmt.Println(departues)

}
