package core

import (
	"encoding/json"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
)

type RequestInterface interface {
	Validate(r *http.Request)
	Parse(r *http.Request)
}

type BaseRequest struct{}

func (BaseRequest) Parse(r *http.Request, rule interface{}) interface{} {
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		xtremelog.Error(err)
		error.ErrXtremeBadRequest(err.Error())
	}

	return rule
}
