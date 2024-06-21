package handler

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
	"service/internal/testing/parser"
	"service/internal/testing/repository"
	"service/internal/testing/service"
)

type TestingHandler struct{}

func (ctr TestingHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewTestingRepository()
	testings, pagination, _ := repo.Find(r.URL.Query())

	psr := parser.TestingParser{Array: testings}

	res := response.Response{Array: psr.Get(), Pagination: &pagination}
	res.Success(w)
}

func (ctr TestingHandler) Create(w http.ResponseWriter, r *http.Request) {
	srv := service.TestingService{}
	srv.Create(w, r)
}

func (ctr TestingHandler) UploadByFile(w http.ResponseWriter, r *http.Request) {
	srv := service.TestingService{}
	srv.UploadByFile(w, r)
}

func (ctr TestingHandler) UploadByContent(w http.ResponseWriter, r *http.Request) {
	srv := service.TestingService{}
	srv.UploadByContent(w, r)
}
