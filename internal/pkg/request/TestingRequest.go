package request

import (
	"github.com/globalxtreme/gobaseconf/middleware"
	"net/http"
	"service/internal/pkg/core"
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
	core.BaseRequest{}.Parse(r, &rule)
}
