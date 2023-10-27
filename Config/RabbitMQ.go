package Config

import (
	"github.com/globalxtreme/gobaseconf/config"
	"os"
	"time"
)

func InitRabbitMQ() {
	config.RabbitMQConf.Connection = config.RabbitMQConnection{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		Username: os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	config.RabbitMQConf.Exchange = config.RabbitMQExchange{
		Name:       "globalxtreme.direct",
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}

	config.RabbitMQConf.Queue = os.Getenv("RABBITMQ_QUEUE")
	config.RabbitMQConf.Timeout = 5 * time.Second
}
