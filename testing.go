package main

import (
	"Service/App/Mail/Testing"
	Testing2 "Service/App/PDF/Testing"
	"Service/Config"
	mail2 "github.com/globalxtreme/gobaseconf/mail"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	tpdf := Testing2.TestingPDF{Name: "Yuswa"}
	path := tpdf.Generate()
	log.Print(path)
}

func sendEmail() {
	Config.InitMail()

	mail := mail2.Mail{}
	mail.Dial(Config.SMTPMail)
	mail.Send(Testing.TestingMail{Name: "Yuswa", Email: "testing1@gmail.com"})
	mail.Send(Testing.TestingMail{Name: "John", Email: "testing2@gmail.com"})
}
