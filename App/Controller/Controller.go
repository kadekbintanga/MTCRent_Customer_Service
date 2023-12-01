package Controller

import (
	"Service/App/Request/Testing"
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

type Controller struct{}

func (ctr Controller) Testing(w http.ResponseWriter, r *http.Request) {
	validate := Testing.TestingRequest{}
	validate.Parse(r)
	validate.Validate(r)

	res := response.Response{}
	res.Success(w)
}
