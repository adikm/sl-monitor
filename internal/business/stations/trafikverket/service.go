package trafikverket

type Service interface {
	FetchStations(authKey string) ([]Station, error)
	FetchDepartures(authKey string) ([]Train, error)
}

type APIService struct {
	remoteClient client
}

func NewAPIService() *APIService {
	return &APIService{remoteClient: &remoteClient{}}
}

var _ Service = &APIService{}
