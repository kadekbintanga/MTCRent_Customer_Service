package Testing

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"github.com/globalxtreme/gobaseconf/middleware"
	"github.com/globalxtreme/gobaseconf/response/error"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type TestingRule struct {
	Name string   `json:"name"`
	Subs []string `json:"subs" validate:"required"`
}

func (rule TestingRule) Validate(r *http.Request) {
	if err := mapstructure.Decode(config.RequestBody, &rule); err != nil {
		xtremelog.Error(err)
		error.ErrXtremeBadRequest(err.Error())
	}

	va := middleware.Validator{}
	va.Make(r, rule)
}
