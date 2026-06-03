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
	form.APIParse(r)
	form.Validate()

	srv := service.NewCustomerService()

	customer := srv.Create(form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	formFilter := form2.CustomerFilterForm{
		UUID: mux.Vars(r)["uuid"],
	}

	srv := service.NewCustomerService()

	srv.Delete(formFilter)
	res := xtremeres.Response{}
	res.Success(w)
}
