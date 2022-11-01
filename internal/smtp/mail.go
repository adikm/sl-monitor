package smtp

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
)

type mailTemplateData struct {
	Station    string
	Date       string
	Canceled   bool
	ShortTrain bool
}

type Mailer struct {
	dialer *gomail.Dialer
	from   string
}

func NewMailer(host string, port int, username, password, from string) *Mailer {
	dialer := gomail.NewDialer(host, port, username, password)

	return &Mailer{dialer, from}
}

func (m *Mailer) SendMail(to string) {
	msg := gomail.NewMessage()

	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Train update")

	var body bytes.Buffer

	t, _ := template.ParseFiles("assets/mail.html")
	t.Execute(&body, mailTemplateData{"TRAIN", "DATA", false, true})

	msg.SetBody("text/html", body.String())

	if err := m.dialer.DialAndSend(msg); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
