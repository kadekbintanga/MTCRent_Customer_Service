package main

import (
	Testing3 "Service/App/Excel/Testing"
	DevTest2 "Service/App/GRPC/Client/DevTest"
	"Service/App/Mail/Testing"
	Testing2 "Service/App/PDF/Testing"
	"Service/Config"
	"Service/RPC/gRPC/DevTest"
	"fmt"
	mail2 "github.com/globalxtreme/gobaseconf/mail"
	"github.com/joho/godotenv"
	"log"
	"time"
)

const (
	address     = "localhost:5050"
	defaultName = "Yuswa"
	timeout     = 5 * time.Second
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.InitRPC()

	fmt.Println(Config.DevTestRPC)

	validation, cleanup := DevTest2.NewValidationClient()
	defer cleanup()

	message, _ := validation.ValidationName("Testing")
	fmt.Println(fmt.Sprintf("Message: %s", message))

	message, _ = validation.ValidationMultiField([]*DevTest.ValidationNameRequest{
		&DevTest.ValidationNameRequest{Name: "John"},
		&DevTest.ValidationNameRequest{Name: "Smith"},
	})
	fmt.Println(fmt.Sprintf("Message: %s", message))
}

func generateExcel() {
	excel := Testing3.TestingExcel{}
	err := excel.Generate()
	if err != nil {
		log.Print(err)
	}
}

func generatePDF() {
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
