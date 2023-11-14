package Config

import (
	"github.com/globalxtreme/gobaseconf/config"
	"os"
)

var (
	SMTPMail config.MailConf
)

func InitMail() {
	SMTPMail = config.MailConf{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}
