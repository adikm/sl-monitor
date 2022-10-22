package trafikverket

func FetchStations(authKey string) ([]Station, error) {
	request := buildStationsRequest(authKey)
	result := new(stationsResult)
	err := post(&request, &result)
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
