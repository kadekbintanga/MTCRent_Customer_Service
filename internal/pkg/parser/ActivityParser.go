package parser

import (
	"service/internal/pkg/model"
	"strings"
)

type ActivityParser struct {
	Activities []model.Activity
	Activity   model.Activity
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
		"createdAt":   activity.CreatedAt.Format("02/01/2006 15:04"),
	}
}
