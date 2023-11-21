package Activity

import (
	"Service/App/Model/Activity"
	"github.com/globalxtreme/gobaseconf/helpers"
	"strings"
)

type ActivityParser struct {
	Activities []Activity.Activity
	Activity   Activity.Activity
}

func (parser ActivityParser) Get() []interface{} {
	var result []interface{}

	for _, activity := range parser.Activities {
		firstParser := ActivityParser{Activity: activity}
		result = append(result, firstParser.First())
	}

	return result
}

func (parser ActivityParser) First() interface{} {
	activity := parser.Activity

	return map[string]interface{}{
		"id":          activity.ID,
		"feature":     strings.ToTitle(strings.ReplaceAll(activity.Feature, "_", " ")),
		"subFeature":  strings.ToTitle(strings.ReplaceAll(activity.SubFeature, "_", " ")),
		"action":      activity.Action,
		"description": activity.Description,
		"reference":   activity.Reference,
		"causedBy":    activity.CausedByName,
		"createdAt":   activity.CreatedAt.Format(helpers.FullDateTimeLayout()),
	}
}
