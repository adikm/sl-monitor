package scheduling

import (
	"bytes"
	"html/template"
	"log"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/business/users"
)

type Sender struct {
	uService  users.Service
	tvService trafikverket.Service
}

func NewSender(uS users.Service, tvService trafikverket.Service) *Sender {
	return &Sender{uS, tvService}
}

func (s *Sender) prepareNotificationMail(userId int) (string, *bytes.Buffer) {

	u, err := s.uService.FindById(userId)
	if err != nil {
		log.Printf("Error while preparing notification mail %d", err)
		return "", nil
	}

	departures, err := s.tvService.FetchDepartures()
	if err != nil {
		return "", nil
	}

	var body bytes.Buffer
	for _, departure := range departures {
		t, _ := template.ParseFiles("assets/mail.html")
		t.Execute(&body, mailTemplateData{u.Name, departure.LineNumber(), s.fullStationName(departure.Destination[0].Code), departure.DepartureTime.String(), departure.Canceled, false})
		break
	}

	return u.Email, &body

}

func (s *Sender) fullStationName(code string) string {
	stations, err := s.tvService.FetchStations()
	if err != nil {
		return code
	}
	for _, station := range stations {
		if station.Code == code {
			return station.Name
		}
	}
	return code
}
