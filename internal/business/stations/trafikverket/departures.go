package trafikverket

import (
	"fmt"
	"log"
)

func (s *APIService) FetchDepartures(stationCode string) ([]Train, error) {
	request := buildDeparturesRequest(stationCode, s.authKey)
	result := new(trainsResult)
	err := s.remoteClient.post(&request, &result)
	if err != nil {
		log.Printf("Error while fetching departures %v\n", err)
		return nil, err
	}
	return result.trains(), nil
}

func buildDeparturesRequest(stationCode, authKey string) request {
	text := fmt.Sprintf(`<OR>
                                            <AND>
                                                <GT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(00:00:10)"/>
                                                <LT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(00:20:00)"/>
                                            </AND>
											<AND>
                                            	<GT name="EstimatedTimeAtLocation" value="$now"/>
											</AND>
                                        </OR>`)
	requestData := request{Login: login{authKey}, Query: query{
		ObjectType:    "TrainAnnouncement",
		SchemaVersion: "1.8",
		OrderBy:       "AdvertisedTimeAtLocation",
		Include:       []string{"ProductInformation", "AdvertisedTimeAtLocation", "EstimatedTimeAtLocation", "Canceled", "Deviation", "ToLocation"},
		Filter: filter{And: and{
			[]equal{
				{Name: "ActivityType", Value: "avgang"},
				{Name: "LocationSignature", Value: stationCode},
			},
			text,
		}},
	}}
	return requestData
}
