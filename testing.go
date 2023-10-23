package main

import (
	"Service/App/Console/Command"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	//Config.InitDB()
	//
	//parameters := url.Values{}
	//parameters.Set("testing", "Test value test")

	cmd := Command.TestCommand{}
	cmd.Handle()
}
