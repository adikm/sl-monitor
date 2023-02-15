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

func (s *Sender) prepareNotificationMail(userId int, stationCode string) (string, *bytes.Buffer) {

	u, err := s.uService.FindById(userId)
	if err != nil {
		log.Printf("Error while preparing notification mail %d", err)
		return "", nil
	}

	departures, err := s.tvService.FetchDepartures(stationCode)
	if err != nil {
		return "", nil
	}

	var body bytes.Buffer
	for _, departure := range departures {
		t, _ := template.ParseFiles("assets/mail.html")
		t.Execute(&body, s.getTemplateData(u.Name, departure))
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

type mailTemplateData struct {
	RecipientName string
	LineNumber    string
	Destination   string
	Date          string
	Canceled      bool
	ShortTrain    bool
}

func (s *Sender) getTemplateData(recipientNam string, departure trafikverket.Train) mailTemplateData {
	return mailTemplateData{
		RecipientName: recipientNam,
		LineNumber:    departure.LineNumber(),
		Destination:   s.fullStationName(departure.Destination[0].Code),
		Date:          departure.DepartureTime.String(),
		Canceled:      departure.Canceled,
		ShortTrain:    false,
	}
}
