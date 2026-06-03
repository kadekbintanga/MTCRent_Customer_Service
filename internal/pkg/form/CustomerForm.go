package form

import (
	"net/http"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"

	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
)

type CustomerForm struct {
	Name            string `json:"name" validate:"required"`
	IDNumber        string `json:"IDNumber" validate:"required,max=100,alphanum"`
	SIMNumber       string `json:"SIMNumber" validate:"required,max=100,alphanum"`
	Phone           string `json:"phone" validate:"required,min=8,max=20,numeric"`
	Address         string `json:"address" validate:"required"`
	StatusId        int    `json:"statusId"`
	BlacklistReason string `json:"blacklistReason"`
}

func (rule *CustomerForm) Validate() {
	va := xtrememdw.Validator{}
	if rule.StatusId != 0 {
		_, exists := constant.CustomerStatus{}.OptionIDNames()[rule.StatusId]
		if !exists {
			error2.ErrXtremeCustomerUpdate("Invalid status")
		} else {
			if rule.StatusId == constant.CUSTOMER_STATUS_BLACKLISTED_ID && rule.BlacklistReason == "" {
				error2.ErrXtremeCustomerUpdate("Blacklist Reason is required")
			}
		}
	}
	va.Make(rule)
}

func (rule *CustomerForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}
