package handler

import (
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/gorilla/mux"

	"service/internal/customer/service"
	form2 "service/internal/pkg/form"
	"service/internal/pkg/parser"
)

type CustomerHandler struct{}

func (ctr CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.CustomerForm{}
	form.APIMultipartParse(r)
	form.Validate()

	srv := service.NewCustomerService()

	customer := srv.Create(form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewCustomerService()

	srv.Delete(mux.Vars(r)["uuid"])
	res := xtremeres.Response{}
	res.Success(w)
}
