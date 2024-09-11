package config

import (
	xtremedb "github.com/globalxtreme/go-core/v2/database"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"os"
	"time"
)

func InitRabbitMQ() {
	xtremerabbitmq.RabbitMQSQL = xtremedb.Connect(xtremedb.DBConf{
		Driver:    xtremedb.MYSQL_DRIVER,
		Host:      os.Getenv("DB_RABBITMQ_HOST"),
		Port:      os.Getenv("DB_RABBITMQ_PORT"),
		Username:  os.Getenv("DB_RABBITMQ_USERNAME"),
		Password:  os.Getenv("DB_RABBITMQ_PASSWORD"),
		Database:  os.Getenv("DB_RABBITMQ_DATABASE"),
		ParseTime: true,
	})

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
