package core

import (
	"encoding/json"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"net/http"
)

type RequestInterface interface {
	Validate(r *http.Request)
	Parse(r *http.Request)
}

type BaseRequest struct{}

func (BaseRequest) Parse(r *http.Request, rule interface{}) interface{} {
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		xtremeres.ErrXtremeBadRequest(err.Error())
	}

	return rule
}
