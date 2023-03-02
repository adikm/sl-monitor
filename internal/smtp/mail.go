package smtp

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"log"
	"sl-monitor/internal/business/users"
)

type Mailer struct {
	*gomail.Dialer
	from         string
	usersService users.Service
}

func NewMailer(host string, port int, username, password, from string, usersService users.Service) *Mailer {
	dialer := gomail.NewDialer(host, port, username, password)

	return &Mailer{dialer, from, usersService}
}

func (m *Mailer) SendMail(to string, body bytes.Buffer) {
	msg := gomail.NewMessage()

	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Train update")

	msg.SetBody("text/html", body.String())

	if err := m.DialAndSend(msg); err != nil {
		log.Printf("Error while sending email %v\n", err)
	}

}
