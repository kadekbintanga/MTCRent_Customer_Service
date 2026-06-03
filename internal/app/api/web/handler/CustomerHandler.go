package handler

import (
	"net/http"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/globalxtreme/go-identifier/data"
	"github.com/gorilla/mux"

	repository2 "service/internal/activity/repository"
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
	form.APIParse(r)
	form.Validate()

	srv := service.NewCustomerService()
	srv.SetEmployeeIdentifier(data.AuthEmployee(r))
	srv.SetActivityRepository(repository2.NewActivityRepository())

	customer := srv.Create(form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	formFilter := form2.CustomerFilterForm{
		UUID: mux.Vars(r)["uuid"],
	}
	form := form2.CustomerForm{}
	form.APIParse(r)
	form.Validate()

	srv := service.NewCustomerService()
	srv.SetEmployeeIdentifier(data.AuthEmployee(r))
	srv.SetActivityRepository(repository2.NewActivityRepository())

	customer := srv.Update(formFilter, form)

	psr := parser.CustomerParser{Object: customer}
	res := xtremeres.Response{Object: psr.First()}
	res.Success(w)
}

func (ctr CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	formFilter := form2.CustomerFilterForm{
		UUID: mux.Vars(r)["uuid"],
	}

	srv := service.NewCustomerService()
	srv.SetEmployeeIdentifier(data.AuthEmployee(r))
	srv.SetActivityRepository(repository2.NewActivityRepository())

	srv.Delete(formFilter)
	res := xtremeres.Response{}
	res.Success(w)
}
