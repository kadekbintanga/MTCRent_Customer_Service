package main

import (
	"Service/Config"
	"Service/Router"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.InitDB()

	newRoute := mux.NewRouter()
	router.RegisterRouter(newRoute, Router.Register)

	config.SetHost()

	fmt.Println(fmt.Sprintf("Server started on %s", config.HostFull))

	err = http.ListenAndServe(config.Host+":"+config.Port, newRoute)
	if err != nil {
		panic(err)
	}
}
