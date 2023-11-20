package main

import (
	"Service/Config"
	"Service/Router"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	config.SetHost()

	Config.InitDB()
	Config.InitCors()
	Config.InitRabbitMQ()
	Config.InitMail()
	Config.InitRPC()

	newCors := cors.New(Config.CorsOptions)

	newRoute := mux.NewRouter()
	router.RegisterRouter(newRoute, Router.Register)

	fmt.Println(fmt.Sprintf("Server started on %s", config.HostFull))

	err = http.ListenAndServe(config.Host+":"+config.Port, newCors.Handler(newRoute))
	if err != nil {
		panic(err)
	}
}
