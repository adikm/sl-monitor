package trafikverket

import (
	"encoding/xml"
	"time"
)

/***
* 𝗥𝗘𝗤𝗨𝗘𝗦𝗧 API model
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
* 𝗥𝗘𝗦𝗣𝗢𝗡𝗦𝗘 API train announcements model
 */

type trainsResult struct {
	Response struct {
		Result []struct {
			Trains []Train `json:"TrainAnnouncement"`
		} `json:"RESULT"`
	} `json:"RESPONSE"`
}
type Train struct {
	DepartureTime time.Time     `json:"AdvertisedTimeAtLocation"`
	Information   []Information `json:"ProductInformation"`
	Canceled      bool          `json:"Canceled"`
	Destination   []Destination `json:"ToLocation"`
	Deviation     []Information `json:"Deviation"`
	//EstimatedTimeAtLocation  time.Time    `json:"EstimatedTimeAtLocation,omitempty"`
}

type Information struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Destination struct {
	Code string `json:"LocationName"`
}

func (t Train) LineNumber() string {
	for _, information := range t.Information {
		if information.Code == "PNA091" {
			return information.Description
		}
	}
	return ""
}

func (t Train) IsShort() bool {
	for _, deviation := range t.Deviation {
		if deviation.Code == "ANA031" {
			return true
		}
	}
	return false
}

func (r trainsResult) trains() []Train {
	return r.Response.Result[0].Trains
}

/**
* 𝗥𝗘𝗦𝗣𝗢𝗡𝗦𝗘 API stations model
 */

type stationsResult struct {
	Response struct {
		Result []struct {
			Stations []Station `json:"TrainStation"`
		} `json:"RESULT"`
	} `json:"RESPONSE"`
}

type Station struct {
	Name string `json:"AdvertisedLocationName"`
	Code string `json:"LocationSignature"`
}

func (r stationsResult) stations() []Station {
	return r.Response.Result[0].Stations
}
