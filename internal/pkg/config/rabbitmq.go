package config

import (
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"os"
	"time"
)

func InitRabbitMQ() {
	xtremerabbitmq.RabbitMQConf.Connection = xtremerabbitmq.RabbitMQConnection{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		Username: os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	xtremerabbitmq.RabbitMQConf.Exchange = xtremerabbitmq.RabbitMQExchange{
		Name:       "globalxtreme.direct",
		Type:       "direct",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}

	xtremerabbitmq.RabbitMQConf.Queue = os.Getenv("RABBITMQ_QUEUE")
	xtremerabbitmq.RabbitMQConf.Timeout = 5 * time.Second
}
