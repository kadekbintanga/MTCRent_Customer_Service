package API

import (
	"Service/App/Controllers"
	"Service/App/Controllers/Web/v1/Activity"
	"Service/App/Controllers/Web/v1/Testing"
	"github.com/gorilla/mux"
)

type Web struct{}

func (api Web) Routers(router *mux.Router) {
	var base Controllers.Controller

	// Default testing router
	router.HandleFunc("/dev-test", base.Testing).Methods("GET")

	// Activity
	router.HandleFunc("/activities", Activity.ActivityController{}.Get).Methods("GET")

	// Testing
	var testingController Testing.TestingController
	router.HandleFunc("/testings", testingController.Get).Methods("GET")
	router.HandleFunc("/testings", testingController.Create).Methods("POST")

}
