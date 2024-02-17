package API

import (
	"Service/App/Handler"
	"Service/App/Handler/Mobile/v1/Testing"
	"github.com/gorilla/mux"
)

type Mobile struct{}

func (api Mobile) Routers(router *mux.Router) {
	var base Handler.Handler

	// Default testing Router
	router.HandleFunc("/dev-test", base.Testing).Methods("GET")

	// Testing
	var testingHandler Testing.TestingHandler
	router.HandleFunc("/testings", testingHandler.Get).Methods("GET")

}
