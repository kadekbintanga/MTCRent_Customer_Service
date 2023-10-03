package Service

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/middleware"
	"net/http"
)

type TestingRule struct {
	Text string `json:"text" validate:"string,nullable"`
}

func (rule TestingRule) Validate(r *http.Request) {
	rule.Text = config.RequestBody["text"].(string)

	validator := middleware.Validator{}
	validator.Make(r, rule)
}
