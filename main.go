package main

import (
	"Service/Config"
	"Service/Router"
	"fmt"
	"github.com/globalxtreme/gobaseconf/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.Init()

	newRoute := mux.NewRouter()
	router.RegisterRouter(newRoute, Router.Register)

	domain := os.Getenv("DOMAIN")
	port := os.Getenv("PORT")

	protocol := "http"
	SSL, _ := strconv.ParseBool(os.Getenv("USE_SSL"))
	if SSL == true {
		protocol = "https"
	}

	fmt.Println(fmt.Sprintf("Server started on %s://%s:%s", protocol, domain, port))

	err = http.ListenAndServe(domain+":"+port, newRoute)
	if err != nil {
		panic(err)
	}
}
