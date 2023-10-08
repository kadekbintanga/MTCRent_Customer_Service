package Controller

import (
	Service "Service/App/Service/Validation/Testing"
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

type Controller struct{}

func (ctr Controller) Testing(w http.ResponseWriter, r *http.Request) {
	validate := Service.TestingRule{}
	validate.Validate(r)

	res := response.Response{}
	res.Success(w)
}
