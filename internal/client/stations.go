package client

import (
	"sl-monitor/internal/config"
)

func FetchStations() []Station {
	request := buildStationsRequest()
	result := new(stationsResult)
	err := post(&request, &result)
	if err != nil {
		return nil
	}
	return result.stations()
}

func buildStationsRequest() request {
	requestData := request{Login: login{config.Cfg.TrafficAPI.AuthKey}, Query: query{
		ObjectType:    "TrainStation",
		SchemaVersion: "1.4",
		Include:       []string{"LocationSignature", "AdvertisedLocationName"},
		Filter: filter{And: and{
			[]equal{
				{Name: "CountryCode", Value: "SE"},
			},
			"",
		}},
	}}
	return requestData
}
