package request

import (
	"github.com/globalxtreme/gobaseconf/middleware"
	"net/http"
	"service/internal/pkg/core"
)

type TestingUploadContentRequest struct {
	Content string `json:"content" validate:"required"`
}

func (rule *TestingUploadContentRequest) Validate(r *http.Request) {
	va := middleware.Validator{}
	va.Make(r, rule)
}

func (rule *TestingUploadContentRequest) Parse(r *http.Request) {
	core.BaseRequest{}.Parse(r, &rule)
}
