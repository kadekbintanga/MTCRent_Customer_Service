package web

import (
	"github.com/gorilla/mux"

	"service/internal/app/api/web/handler"
)

func Register(router *mux.Router) {
	activityRouter(router)
	testingRouter(router) // TODO: Hanya contoh. nanti langsung hapus saja
	customerRouter(router)
}

func activityRouter(router *mux.Router) {
	router.HandleFunc("/activities", handler.ActivityHandler{}.Get).Methods("GET")
}

// TODO: Hanya contoh. nanti langsung hapus saja
func testingRouter(router *mux.Router) {
	var testingHandler handler.TestingHandler
	router.HandleFunc("/testings", testingHandler.Get).Methods("GET")
	router.HandleFunc("/testings", testingHandler.Create).Methods("POST")
	router.HandleFunc("/testings/upload/file", testingHandler.UploadByFile).Methods("POST")
	router.HandleFunc("/testings/upload/content", testingHandler.UploadByContent).Methods("POST")
}

func customerRouter(router *mux.Router) {
	var staticHandler handler.CustomerStaticHandler
	router.HandleFunc("/components/statics/customer-statuses", staticHandler.CustomerStatus).Methods("GET")

	var customerHandler handler.CustomerHandler
	router.HandleFunc("", customerHandler.Get).Methods("GET")
	router.HandleFunc("", customerHandler.Create).Methods("POST")
	router.HandleFunc("/{uuid}", customerHandler.Detail).Methods("GET")
	router.HandleFunc("/{uuid}", customerHandler.Update).Methods("PUT")
	router.HandleFunc("/{uuid}", customerHandler.Delete).Methods("DELETE")
}
