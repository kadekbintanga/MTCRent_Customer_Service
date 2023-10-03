package Activity

import (
	"Service/App/Models/Testing"
	"github.com/globalxtreme/gobaseconf/helpers"
)

type TestingParser struct {
	Array  []Testing.Testing
	Object Testing.Testing
}

func (parser TestingParser) Get() []interface{} {
	var result []interface{}

	for _, activity := range parser.Array {
		firstParser := TestingParser{Object: activity}
		result = append(result, firstParser.First())
	}

	return result
}

func (parser TestingParser) First() interface{} {
	activity := parser.Object

	return map[string]interface{}{
		"id":        activity.ID.ID,
		"name":      activity.Name,
		"createdAt": activity.CreatedAt.Format(helpers.FullDateTimeLayout()),
	}
}

func (parser TestingParser) CreateActivity(action string) interface{} {
	return parser.First()
}

func (parser TestingParser) UpdateActivity(action string) interface{} {
	return parser.First()
}

func (parser TestingParser) DeleteActivity(action string) interface{} {
	return parser.First()
}

func (parser TestingParser) GeneralActivity(action string) interface{} {
	if action == "onlyName" {
		activity := parser.Object

		return map[string]interface{}{
			"name": activity.Name,
		}
	}

	return parser.First()
}
