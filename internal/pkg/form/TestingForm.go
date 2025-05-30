package form

import (
	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
	"net/http"
	"service/internal/pkg/core"
)

type TestingForm struct {
	Name string   `json:"name"`
	Subs []string `json:"subs" validate:"required"`
}

func (rule *TestingForm) Validate(r *http.Request) {
	va := xtrememdw.Validator{}
	va.Make(r, rule)
}

func (rule *TestingForm) Parse(r *http.Request) {
	core.BaseRequest{}.Parse(r, &rule)
}
