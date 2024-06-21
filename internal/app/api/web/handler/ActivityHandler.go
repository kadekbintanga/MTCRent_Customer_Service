package handler

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
	parser2 "service/internal/activity/parser"
	"service/internal/activity/repository"
)

type ActivityHandler struct{}

func (ctr ActivityHandler) Get(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewActivityRepository()
	activities, pagination, _ := repo.Find(r.URL.Query())

	parser := parser2.ActivityParser{Activities: activities}

	res := response.Response{Array: parser.Get(), Pagination: &pagination}
	res.Success(w)
}
