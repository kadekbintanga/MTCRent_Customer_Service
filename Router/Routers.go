package Router

import (
	"Service/Router/API"
	"fmt"
	"github.com/globalxtreme/go-identifier/middleware"
	"github.com/gorilla/mux"
	"os"
)

func Register(router *mux.Router) {
	version := os.Getenv("VERSION")
	service := os.Getenv("SERVICE")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.EmployeeIdentifier)

	web := API.Web{}
	web.Routers(api.PathPrefix(fmt.Sprintf("/web/%s/%s", version, service)).Subrouter())

	mobile := API.Mobile{}
	mobile.Routers(api.PathPrefix(fmt.Sprintf("/mobile/%s/%s", version, service)).Subrouter())
}
