package client

import (
	"fmt"
	"sl-monitor/internal/config"
)

func FetchStations() {
	request := buildStationsRequest()
	result := new(stationsResult)
	post(&request, &result)
	fmt.Println(result)
}

func buildStationsRequest() request {
	requestData := request{Login: login{config.Cfg.TrafficAPI.AuthKey}, Query: query{
		ObjectType:    "TrainStation",
		SchemaVersion: "1.4",
		Include:       []string{"LocationSignature", "AdvertisedLocationName", "Prognosticated"},
	}}
	return requestData
}
