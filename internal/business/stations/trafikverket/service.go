package trafikverket

type Service interface {
	FetchStations(authKey string) ([]Station, error)
	FetchDepartures(authKey string) []Train
}

type APIService struct {
	remoteClient client
}

func NewAPIService() *APIService {
	return &APIService{remoteClient: &remoteClient{}}
}
