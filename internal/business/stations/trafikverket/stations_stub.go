package trafikverket

type ServiceStub struct {
}

func (s *ServiceStub) FetchStations() ([]Station, error) {
	return []Station{}, nil
}
func (s *ServiceStub) FetchDepartures() ([]Train, error) {
	return []Train{}, nil
}

var _ Service = &ServiceStub{}
