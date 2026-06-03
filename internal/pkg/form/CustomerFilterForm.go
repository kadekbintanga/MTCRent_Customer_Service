package form

import (
	"net/url"

	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
)

type CustomerFilterForm struct {
	ID             uint
	UUID           string
	IDNumber       string
	SIMNumber      string
	Search         string
	Orders         map[string]string
	WithPagination bool
	Page           int
	Limit          int
}

func (f *CustomerFilterForm) FilterParse(parameter url.Values) {
	if searchReq := parameter.Get("search"); len(searchReq) >= 3 {
		f.Search = searchReq
	}

	f.WithPagination = true
	if pageReq := parameter.Get("page"); pageReq != "" {
		f.Page = xtremepkg.ToInt(pageReq)
		f.Limit = xtremepkg.ToInt(parameter.Get("limit"))
	} else {
		f.Page = 1
	}
}
