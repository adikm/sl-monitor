package trafikverket

type ServiceStub struct {
}

func (s *ServiceStub) FetchStations(authKey string) ([]Station, error) {
	return []Station{}, nil
}
func (s *ServiceStub) FetchDepartures(authKey string) ([]Train, error) {
	return []Train{}, nil
}

var _ Service = &ServiceStub{}
