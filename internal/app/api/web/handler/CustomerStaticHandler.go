package handler

import (
	"net/http"
	"service/internal/pkg/constant"
	"service/internal/pkg/core"

	xtremeres "github.com/globalxtreme/go-core/v2/response"
)

type CustomerStaticHandler struct{}

func (ctr CustomerStaticHandler) CustomerStatus(w http.ResponseWriter, r *http.Request) {
	customerStatus := core.IDName{}.Get(constant.CustomerStatus{})

	res := xtremeres.Response{Array: customerStatus}
	res.Success(w)
}
