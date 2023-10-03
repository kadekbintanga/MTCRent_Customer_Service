package main

import (
	"Service/Config"
	"Service/Database/Migrations"
	"github.com/globalxtreme/gobaseconf/database/migration"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.Init()
	migration.Migrate(Migrations.Tables(), Migrations.Columns())
}
