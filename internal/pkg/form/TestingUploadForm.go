package form

import (
	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"
	"net/http"
)

type TestingUploadForm struct {
	Request *http.Request
}

func (rule *TestingUploadForm) Validate(r *http.Request) {
	va := xtrememdw.Validator{}
	va.Make(r, rule)
}

func (rule *TestingUploadForm) Parse(r *http.Request) {
}
