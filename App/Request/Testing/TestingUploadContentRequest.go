package Testing

import (
	"Service/App/Request"
	"github.com/globalxtreme/gobaseconf/middleware"
	"net/http"
)

type TestingUploadContentRequest struct {
	Content string `json:"content" validate:"required"`
}

func (rule *TestingUploadContentRequest) Validate(r *http.Request) {
	va := middleware.Validator{}
	va.Make(r, rule)
}

func (rule *TestingUploadContentRequest) Parse(r *http.Request) {
	Request.BaseRequest{}.Parse(r, &rule)
}
