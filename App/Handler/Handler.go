package Handler

import (
	"Service/App/Request/Testing"
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

type Handler struct{}

func (ctr Handler) Testing(w http.ResponseWriter, r *http.Request) {
	validate := Testing.TestingRequest{}
	validate.Parse(r)
	validate.Validate(r)

	res := response.Response{}
	res.Success(w)
}
