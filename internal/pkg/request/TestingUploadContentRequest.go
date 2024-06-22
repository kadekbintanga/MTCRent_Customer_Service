package request

import (
	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
	"net/http"
	"service/internal/pkg/core"
)

type TestingUploadContentRequest struct {
	Content string `json:"content" validate:"required"`
}

func (rule *TestingUploadContentRequest) Validate(r *http.Request) {
	va := xtrememdw.Validator{}
	va.Make(r, rule)
}

func (rule *TestingUploadContentRequest) Parse(r *http.Request) {
	core.BaseRequest{}.Parse(r, &rule)
}
