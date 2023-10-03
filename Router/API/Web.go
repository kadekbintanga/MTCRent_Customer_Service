package API

import (
	"Service/App/Controllers"
	"github.com/gorilla/mux"
)

type Web struct{}

func (api Web) Routers(router *mux.Router) {
	var base Controllers.Controller

	// Default testing router
	router.HandleFunc("/dev-test", base.Testing).Methods("GET")

}
