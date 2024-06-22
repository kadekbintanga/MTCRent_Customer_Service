package config

import (
	xtremedb "github.com/globalxtreme/go-core/v2/database"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"gorm.io/gorm"
	"os"
)

var (
	PgSQL *gorm.DB
)

func InitDB() {
	PgSQL = xtremedb.Connect(xtremedb.DBConf{
		Driver:    xtremedb.POSTGRESQL_DRIVER,
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		ParseTime: true,
	})

	xtremerabbitmq.RabbitMQSQL = xtremedb.Connect(xtremedb.DBConf{
		Driver:    xtremedb.MYSQL_DRIVER,
		Host:      os.Getenv("DB_RABBITMQ_HOST"),
		Port:      os.Getenv("DB_RABBITMQ_PORT"),
		Username:  os.Getenv("DB_RABBITMQ_USERNAME"),
		Password:  os.Getenv("DB_RABBITMQ_PASSWORD"),
		Database:  os.Getenv("DB_RABBITMQ_DATABASE"),
		ParseTime: true,
	})
}
