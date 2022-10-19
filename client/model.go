package client

import (
	"encoding/xml"
	"time"
)

/***
* 𝗥𝗘𝗤𝗨𝗘𝗦𝗧 model
 */
type request struct {
	XMLName xml.Name `xml:"REQUEST"`
	Login   login    `xml:"LOGIN"`
	Query   query    `xml:"QUERY"`
}

type login struct {
	AuthenticationKey string `xml:"authenticationkey,attr"`
}

type query struct {
	ObjectType    string   `xml:"objecttype,attr"`
	SchemaVersion string   `xml:"schemaversion,attr"`
	OrderBy       string   `xml:"orderby,attr"`
	Filter        filter   `xml:"FILTER"`
	Include       []string `xml:"INCLUDE"`
}

type filter struct {
	And and `xml:"AND"`
}
type equal struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
type and struct {
	Equals []equal `xml:"EQ"`
	Text   string  `xml:",innerxml"`
}

/**
* 𝗥𝗘𝗦𝗣𝗢𝗡𝗦𝗘 model
 */

type requestResult struct {
	Response struct {
		Result []struct {
			Trains []Train `json:"TrainAnnouncement"`
		} `json:"RESULT"`
	} `json:"RESPONSE"`
}
type Train struct {
	DepartureTime time.Time     `json:"AdvertisedTimeAtLocation"`
	Canceled      bool          `json:"Canceled"`
	TrainId       string        `json:"TechnicalTrainIdent"`
	Destination   []Destination `json:"ToLocation"`
	//EstimatedTimeAtLocation  time.Time    `json:"EstimatedTimeAtLocation,omitempty"`
}
type Destination struct {
	Name string `json:"LocationName"`
}

func (r requestResult) trains() []Train {
	return r.Response.Result[0].Trains
}
