package config

import (
	xtremedb "github.com/globalxtreme/go-core/v2/database"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	PgSQL *gorm.DB
)

func InitDB() func() {
	PgSQL = xtremedb.Connect(xtremedb.DBConf{
		Driver:    xtremedb.POSTGRESQL_DRIVER,
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		ParseTime: true,
	})

	pgsqlDB, err := PgSQL.DB()
	if err != nil {
		log.Panicf("Getting DB object is failed: %s", err.Error())
	}

	closeDB := func() {
		pgsqlDB.Close()
	}

	return closeDB
}
