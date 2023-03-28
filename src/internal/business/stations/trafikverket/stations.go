package trafikverket

import (
	"time"
)

var cachedStation struct {
	stations []Station
	updated  time.Time
}

// FetchStations fetches and caches result for 24 hours
func (s *APIService) FetchStations() ([]Station, error) {
	accessedMoreThan24Hours := time.Now().Sub(cachedStation.updated).Hours() > 24
	if accessedMoreThan24Hours {
		request := buildStationsRequest(s.authKey)
		result := new(stationsResult)
		err := s.remoteClient.post(&request, &result)
		if err != nil {
			return nil, err
		}
		cachedStation.stations = result.stations()
		cachedStation.updated = time.Now()
	}

	return cachedStation.stations, nil
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
