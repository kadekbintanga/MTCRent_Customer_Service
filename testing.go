package main

import (
	"Service/Config"
	"fmt"
	"github.com/joho/godotenv"
	"net/url"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.Init()

	parameters := url.Values{}
	parameters.Set("testing", "Test value test")

	fmt.Println(parameters)
}
