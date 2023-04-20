package trafikverket

import "sl-monitor/internal/cache"

type Service interface {
	FetchStations() ([]Station, error)
	FetchDepartures(stationCode string) ([]Train, error)
}

type APIService struct {
	remoteClient client
	cache        cache.Client
	authKey      string
}

func NewAPIService(cache cache.Client, authKey string) *APIService {
	return &APIService{&remoteClient{}, cache, authKey}
}

var _ Service = &APIService{}
