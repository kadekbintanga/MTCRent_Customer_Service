package Testing

import (
	Testing2 "Service/App/Algorithm/Testing"
	"Service/App/Model/Testing"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"log"
	"net/http"
)

type TestingController struct{}

func (ctr TestingController) Get(w http.ResponseWriter, r *http.Request) {
	testing := Testing.Testing{}
	err := Config.PgSQL.Where("testing", "asdf").First(&testing).Error
	xtremelog.Error(err)
	log.Panicf("Testing %s", err)
	//repo := TestingRepo.TestingRepository{}
	//testings, pagination, _ := repo.Get(r.URL.Query())

	//parser := TestingParser.TestingParser{Array: testings}
	//
	//res := response.Response{Array: parser.Get(), Pagination: &pagination}
	//res.Success(w)
}

func (ctr TestingController) Create(w http.ResponseWriter, r *http.Request) {
	algo := Testing2.TestingAlgo{}
	algo.Create(w, r)
}

func (ctr TestingController) UploadByFile(w http.ResponseWriter, r *http.Request) {
	algo := Testing2.TestingAlgo{}
	algo.UploadByFile(w, r)
}

func (ctr TestingController) UploadByContent(w http.ResponseWriter, r *http.Request) {
	algo := Testing2.TestingAlgo{}
	algo.UploadByContent(w, r)
}
