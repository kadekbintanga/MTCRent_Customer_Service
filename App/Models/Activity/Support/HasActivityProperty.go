package Support

import (
	"Service/App/Models/Testing"
	Activity "Service/App/Parser/Testing"
	"Service/App/Services/General"
)

type HasActivityProperty struct {
	Base General.BaseActivityPropertyParserInterface
}

func (property *HasActivityProperty) SetParser(model interface{}) *HasActivityProperty {
	switch model.(type) {
	case Testing.Testing:
		property.Base = &Activity.TestingParser{Object: model.(Testing.Testing)}
	}

	return property
}
