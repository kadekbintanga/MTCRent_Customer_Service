package Testing

import (
	"Service/App/Request"
	"github.com/globalxtreme/gobaseconf/middleware"
	"net/http"
)

type TestingRequest struct {
	Name string   `json:"name"`
	Subs []string `json:"subs" validate:"required"`
}

func (rule *TestingRequest) Validate(r *http.Request) {
	va := middleware.Validator{}
	va.Make(r, rule)
}

func (rule *TestingRequest) Parse(r *http.Request) {
	Request.BaseRequest{}.Parse(r, &rule)
}
