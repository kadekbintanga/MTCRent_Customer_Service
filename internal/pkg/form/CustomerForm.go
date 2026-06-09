package form

import (
	"net/http"

	xtrememdw "github.com/globalxtreme/go-core/v2/middleware"

	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
)

type CustomerPhotoForm struct {
	Link     *string `form:"link"`
	MimeType *string `form:"mimeType" validate:"required"`
	Deleted  *bool   `json:"deleted" form:"deleted"`
}

type CustomerForm struct {
	Request         *http.Request
	Name            string             `form:"name" validate:"required"`
	IDNumber        string             `form:"IDNumber" validate:"required,max=100,alphanum"`
	SIMNumber       string             `form:"SIMNumber" validate:"required,max=100,alphanum"`
	Phone           string             `form:"phone" validate:"required,min=8,max=20,numeric"`
	Address         string             `form:"address" validate:"required"`
	StatusId        int                `form:"statusId" validate:"required"`
	BlacklistReason string             `form:"blacklistReason"`
	IdentityPhoto   *CustomerPhotoForm `form:"identityPhoto"`
}

func (rule *CustomerForm) Validate() {
	va := xtrememdw.Validator{}
	_, exists := constant.CustomerStatus{}.OptionIDNames()[rule.StatusId]
	if !exists {
		error2.ErrXtremeCustomerUpdate("Invalid status")
	} else {
		if rule.StatusId == constant.CUSTOMER_STATUS_BLACKLISTED_ID && rule.BlacklistReason == "" {
			error2.ErrXtremeCustomerUpdate("Blacklist Reason is required")
		}
	}
	va.Make(rule)
}

func (rule *CustomerForm) APIParse(r *http.Request) {
	core.BaseForm{}.APIParse(r, &rule)
}

func (rule *CustomerForm) APIMultipartParse(r *http.Request) {
	rule.Request = r
	core.BaseForm{}.APIMultipartParse(r, &rule)
}
