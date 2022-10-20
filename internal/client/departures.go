package client

import (
	"fmt"
	"sl-monitor/internal/config"
)

func FetchDepartures() {
	request := buildDeparturesRequest()
	result := new(trainsResult)
	post(&request, &result)
	fmt.Println(result.trains())
}

func buildDeparturesRequest() request {
	text := fmt.Sprintf(`<OR>
                                            <AND>
                                                <GT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(-00:15:00)"/>
                                                <LT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(00:15:00)"/>
                                            </AND>
                                            <GT name="EstimatedTimeAtLocation" value="$now"/>
                                        </OR>`)
	requestData := request{Login: login{config.Cfg.TrafficAPI.AuthKey}, Query: query{
		ObjectType:    "TrainAnnouncement",
		SchemaVersion: "1.6",
		OrderBy:       "AdvertisedTimeAtLocation",
		Include:       []string{"TechnicalTrainIdent", "AdvertisedTimeAtLocation", "EstimatedTimeAtLocation", "Canceled", "Deviation", "ToLocation"},
		Filter: filter{And: and{
			[]equal{{
				Name:  "ActivityType",
				Value: "avgang"}, {Name: "LocationSignature", Value: "Hnd"},
			},
			text,
		}},
	}}
	return requestData
}
