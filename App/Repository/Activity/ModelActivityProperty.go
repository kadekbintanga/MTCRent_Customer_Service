package Activity

import (
	"Service/App/Model/Testing"
	Activity "Service/App/Parser/Testing"
	"Service/App/Service/General"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
)

type ModelActivityProperty struct {
	Base General.BaseActivityPropertyParserInterface
}

func (property *ModelActivityProperty) SetParser(model interface{}) *ModelActivityProperty {
	switch model.(type) {
	case Testing.Testing:
		property.Base = &Activity.TestingParser{Object: model.(Testing.Testing)}
		break
	default:
		xtremelog.Error("ACTIVITY-ERROR: Model not registered")
		break
	}

	return property
}
