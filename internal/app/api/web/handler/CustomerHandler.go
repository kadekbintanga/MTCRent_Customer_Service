package handler

import (
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/globalxtreme/go-identifier/data"
	"github.com/gorilla/mux"

	"service/internal/customer/repository"
	"service/internal/customer/service"
	form2 "service/internal/pkg/form"
	"service/internal/pkg/parser"
)

type CustomerHandler struct{}

func (ctr CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	form := form2.CustomerFilterForm{
		Orders: map[string]string{"name": "ASC"},
	}

	form.FilterParse(r.URL.Query())

	repo := repository.NewCustomerRepository()
	customers, pagination := repo.PaginateByForm(form)

	psr := parser.CustomerParser{Array: customers}

	res := xtremeres.Response{Array: psr.Get(), Pagination: &pagination}
	res.Success(w)
}

func (ctr CustomerHandler) Detail(w http.ResponseWriter, r *http.Request) {
	form := form2.CustomerFilterForm{
		UUID: mux.Vars(r)["uuid"],
	}

	repo := repository.NewCustomerRepository()
	customer := repo.FirstByForm(form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	form := form2.CustomerForm{}
	form.APIMultipartParse(r)
	form.Validate()

	srv := service.NewCustomerService()
	srv.SetEmployeeIdentifier(data.AuthEmployee(r))
	customer := srv.Create(form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	form := form2.CustomerForm{}
	form.APIMultipartParse(r)
	form.Validate()

	srv := service.NewCustomerService()
	srv.SetEmployeeIdentifier(data.AuthEmployee(r))

	customer := srv.Update(mux.Vars(r)["uuid"], form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	srv := service.NewCustomerService()
	srv.SetEmployeeIdentifier(data.AuthEmployee(r))

	srv.Delete(mux.Vars(r)["uuid"])
	res := xtremeres.Response{}
	res.Success(w)
}
