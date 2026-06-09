package form

import (
	"net/http"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"

	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
)

type CustomerStatusForm struct {
	StatusId        int    `json:"statusId" validate:"required"`
	BlacklistReason string `json:"blacklistReason"`
	CreatedBy       string `form:"createdByUUID"`
	CreatedByName   string `form:"createdByName"`
}

func (rule *CustomerStatusForm) Validate() {
	va := xtrememdw.Validator{}
	_, status := constant.CustomerStatus{}.OptionIDNames()[rule.StatusId]
	if !status {
		error2.ErrXtremeCustomerUpdate("Invalid status")
	} else {
		if rule.StatusId == constant.CUSTOMER_STATUS_BLACKLISTED_ID && rule.BlacklistReason == "" {
			error2.ErrXtremeCustomerUpdate("Blacklist Reason is required")
		}
	}
	va.Make(rule)
}

func (rule *CustomerStatusForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}

func (rule *CustomerStatusForm) AsyncWorkflowParse(payload interface{}) error {
	return core.BaseForm{}.AsyncWorkflowParse(payload, &rule)
}
