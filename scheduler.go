package main

import (
	"Service/App/Console"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/console"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.InitDB()

	console.Schedules(Console.Schedules)
}
