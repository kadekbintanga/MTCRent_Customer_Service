package API

import (
	"Service/App/Controller"
	"Service/App/Controller/Mobile/v1/Testing"
	"github.com/gorilla/mux"
)

type Mobile struct{}

func (api Mobile) Routers(router *mux.Router) {
	var base Controller.Controller

	// Default testing router
	router.HandleFunc("/dev-test", base.Testing).Methods("GET")

	// Testing
	var testingController Testing.TestingController
	router.HandleFunc("/testings", testingController.Get).Methods("GET")

}
