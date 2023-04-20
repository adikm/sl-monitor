package trafikverket

import (
	"encoding/json"
	"sl-monitor/internal/cache"
	"time"
)

var cachedStation struct {
	stations []Station
	updated  time.Time
}

// FetchStations fetches and caches result for 24 hours
func (s *APIService) FetchStations() ([]Station, error) {
	value := cache.Instance.FetchValue("stations")
	result := new(stationsResult)
	if value == "" {
		request := buildStationsRequest(s.authKey)
		err := s.remoteClient.post(&request, &result)
		if err != nil {
			return nil, err
		}
		stations := result.stations()
		marshal, _ := json.Marshal(stations)
		cache.Instance.SetValue("stations", string(marshal), 24*time.Hour)
		return stations, nil
	}

	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}
	return result.stations(), nil
}

func buildStationsRequest(authKey string) request {
	requestData := request{Login: login{authKey}, Query: query{
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
