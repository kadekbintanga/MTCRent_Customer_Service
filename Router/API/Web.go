package API

import (
	"Service/App/Controller"
	"Service/App/Controller/Web/v1/Activity"
	"Service/App/Controller/Web/v1/Testing"
	"github.com/gorilla/mux"
)

type Web struct{}

func (api Web) Routers(router *mux.Router) {
	var base Controller.Controller

	// Default testing Router
	router.HandleFunc("/dev-test", base.Testing).Methods("GET")

	// Activity
	router.HandleFunc("/activities", Activity.ActivityController{}.Get).Methods("GET")

	// Testing
	var testingController Testing.TestingController
	router.HandleFunc("/testings", testingController.Get).Methods("GET")
	router.HandleFunc("/testings", testingController.Create).Methods("POST")
	router.HandleFunc("/testings/upload/file", testingController.UploadByFile).Methods("POST")
	router.HandleFunc("/testings/upload/content", testingController.UploadByContent).Methods("POST")

}
