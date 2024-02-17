package API

import (
	"Service/App/Handler"
	"Service/App/Handler/Web/v1/Activity"
	"Service/App/Handler/Web/v1/Testing"
	"github.com/gorilla/mux"
)

type Web struct{}

func (api Web) Routers(router *mux.Router) {
	var base Handler.Handler

	// Default testing Router
	router.HandleFunc("/dev-test", base.Testing).Methods("GET")

	// Activity
	router.HandleFunc("/activities", Activity.ActivityHandler{}.Get).Methods("GET")

	// Testing
	var testingHandler Testing.TestingHandler
	router.HandleFunc("/testings", testingHandler.Get).Methods("GET")
	router.HandleFunc("/testings", testingHandler.Create).Methods("POST")
	router.HandleFunc("/testings/upload/file", testingHandler.UploadByFile).Methods("POST")
	router.HandleFunc("/testings/upload/content", testingHandler.UploadByContent).Methods("POST")

}
