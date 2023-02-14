package trafikverket

type Service interface {
	FetchStations() ([]Station, error)
	FetchDepartures() ([]Train, error)
}

type APIService struct {
	remoteClient client
	authKey      string
}

func NewAPIService(authKey string) *APIService {
	return &APIService{remoteClient: &remoteClient{}, authKey: authKey}
}

var _ Service = &APIService{}
