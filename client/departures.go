package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sl-monitor/config"
)

func FetchDepartures() {
	request := buildRequest()
	xmlRequestData, err := xml.Marshal(request)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(config.Cfg.TrafficAPI.URL, "text/xml", bytes.NewBuffer(xmlRequestData))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var result requestResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result.trains())
}

func buildRequest() request {
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
