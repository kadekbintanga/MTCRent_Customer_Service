package main

import (
	"Service/App/Mail"
	"Service/Config"
	mail2 "github.com/globalxtreme/gobaseconf/mail"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.InitMail()

	mail := mail2.Mail{}
	mail.Dial(Config.SMTPMail).
		Send(Mail.TestingMail{Name: "Yuswa", Email: "testing1@gmail.com"}).
		Send(Mail.TestingMail{Name: "John", Email: "testing2@gmail.com"})
}
