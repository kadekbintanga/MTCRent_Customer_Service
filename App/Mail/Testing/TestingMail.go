package Testing

import (
	"github.com/globalxtreme/gobaseconf/helpers"
	"gopkg.in/mail.v2"
)

type TestingMail struct {
	Name  string
	Email string
}

func (m TestingMail) Message() *mail.Message {
	msg := mail.NewMessage()
	msg.SetHeader("To", m.Email)
	msg.SetHeader("Subject", "Hello First!")
	msg.SetBody("text/html", helpers.MailHTMLTemplate("EmailTemplate.html", m))

	return msg
}
