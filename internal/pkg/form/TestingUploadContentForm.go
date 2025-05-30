package form

import (
	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
	"net/http"
	"service/internal/pkg/core"
)

type TestingUploadContentForm struct {
	Content string `json:"content" validate:"required"`
}

func (rule *TestingUploadContentForm) Validate(r *http.Request) {
	va := xtrememdw.Validator{}
	va.Make(r, rule)
}

func (rule *TestingUploadContentForm) Parse(r *http.Request) {
	core.BaseRequest{}.Parse(r, &rule)
}
