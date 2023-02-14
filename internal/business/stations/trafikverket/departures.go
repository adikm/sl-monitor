package trafikverket

import (
	"fmt"
)

func (s *APIService) FetchDepartures() ([]Train, error) {
	request := buildDeparturesRequest(s.authKey)
	result := new(trainsResult)
	err := s.remoteClient.post(&request, &result)
	if err != nil {
		return nil, err
	}
	return result.trains(), nil
}

func buildDeparturesRequest(authKey string) request {
	text := fmt.Sprintf(`<OR>
                                            <AND>
                                                <GT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(-00:15:00)"/>
                                                <LT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(00:15:00)"/>
                                            </AND>
                                            <GT name="EstimatedTimeAtLocation" value="$now"/>
                                        </OR>`)
	requestData := request{Login: login{authKey}, Query: query{
		ObjectType:    "TrainAnnouncement",
		SchemaVersion: "1.8",
		OrderBy:       "AdvertisedTimeAtLocation",
		Include:       []string{"ProductInformation", "AdvertisedTimeAtLocation", "EstimatedTimeAtLocation", "Canceled", "Deviation", "ToLocation"},
		Filter: filter{And: and{
			[]equal{
				{Name: "ActivityType", Value: "avgang"},
				{Name: "LocationSignature", Value: "Hnd"},
			},
			text,
		}},
	}}
	return requestData
}
