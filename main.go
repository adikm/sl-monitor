package main

import (
	"encoding/xml"
	"fmt"
)

type Request struct {
	XMLName xml.Name `xml:"REQUEST"`
	Login   Login    `xml:"LOGIN"`
	Query   Query    `xml:"QUERY"`
}

type Login struct {
	AuthenticationKey string `xml:"authenticationkey,attr"`
}

type Query struct {
	ObjectType    string   `xml:"objecttype,attr"`
	SchemaVersion string   `xml:"schemaversion,attr"`
	OrderBy       string   `xml:"orderby,attr"`
	Filter        Filter   `xml:"FILTER"`
	Include       []string `xml:"INCLUDE"`
}

type Filter struct {
	And And `xml:"AND"`
}
type Equal struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
type And struct {
	Equals []Equal `xml:"EQ"`
	Text   string  `xml:",innerxml"`
}

func main() {
	//currentTime := time.Now()
	//minusHour := currentTime.Add(-1 * time.Hour)
	//plusHour := currentTime.Add(1 * time.Hour)
	text := fmt.Sprintf(`<OR>
                                            <AND>
                                                <GT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(-00:15:00)"/>
                                                <LT name="AdvertisedTimeAtLocation"
                                                            value="$dateadd(00:15:00)"/>
                                            </AND>
                                            <GT name="EstimatedTimeAtLocation" value="$now"/>
                                        </OR>`)
	request := Request{Login: Login{AuthenticationKey: "0e7862ebcacf4d7a90c2a90a443bca3f"},
		Query: Query{
			ObjectType:    "TrainAnnouncement",
			SchemaVersion: "1.6",
			OrderBy:       "AdvertisedTimeAtLocation",
			Include:       []string{"TechnicalTrainIdent", "AdvertisedTimeAtLocation", "EstimatedTimeAtLocation", "Canceled", "Deviation", "ToLocation"},
			Filter: Filter{And: And{Equals: []Equal{{
				Name:  "ActivityType",
				Value: "avgang"}, {Name: "LocationSignature", Value: "Hnd"},
			},
				Text: text}},
		},
	}

	indent, err := xml.MarshalIndent(request, "", " ")
	if err == nil {
		fmt.Println(string(indent))
	}

}
