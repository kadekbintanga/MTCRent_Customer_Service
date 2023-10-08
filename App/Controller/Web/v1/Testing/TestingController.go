package Testing

import (
	Testing2 "Service/App/Algorithm/Testing"
	TestingParser "Service/App/Parser/Testing"
	TestingRepo "Service/App/Repository/Testing"
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

type TestingController struct{}

func (ctr TestingController) Get(w http.ResponseWriter, r *http.Request) {
	repo := TestingRepo.TestingRepository{}
	testings, pagination, _ := repo.Get(r.URL.Query())

	parser := TestingParser.TestingParser{Array: testings}

	res := response.Response{Array: parser.Get(), Pagination: &pagination}
	res.Success(w)
}

func (ctr TestingController) Create(w http.ResponseWriter, r *http.Request) {
	algo := Testing2.TestingAlgo{}
	algo.Create(w, r)
}
