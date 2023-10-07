package Activity

import (
	"Service/App/Models/Activity"
	Server "Service/App/Services/Constant/Activity"
	"Service/App/Services/Error"
	"Service/App/Services/General"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/data"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
)

type property struct {
	Old interface{}
	New interface{}
}

type UseActivity struct {
	Feature     string                         `gorm:"-"`
	SubFeature  string                         `gorm:"-"`
	Action      string                         `gorm:"-"`
	Description string                         `gorm:"-"`
	Property    property                       `gorm:"-"`
	Model       General.ActivityModelInterface `gorm:"-"`
}

func (aa UseActivity) SetSubFeature(subFeature string) UseActivity {
	aa.SubFeature = subFeature

	return aa
}

func (aa UseActivity) SetModel(model General.ActivityModelInterface) UseActivity {
	aa.Model = model

	return aa
}

func (aa UseActivity) SetOldProperty(action string, subs ...string) UseActivity {
	aa.Action = action

	if aa.Model != nil {
		aa.Property.Old = aa.setPropertyWithParser(action, subs...)
	}

	return aa
}

func (aa UseActivity) SetNewProperty(action string, subs ...string) UseActivity {
	aa.Action = action

	if aa.Model != nil {
		aa.Property.New = aa.setPropertyWithParser(action, subs...)
	}

	return aa
}

func (aa UseActivity) Save(description string) error {
	var activity Activity.Activity
	activity.Feature = aa.Model.TableName()
	activity.SubFeature = aa.SubFeature
	activity.Action = aa.Action
	activity.Description = description
	activity.Reference = aa.Model.SetReference()
	activity.Properties = map[string]interface{}{
		"old": aa.Property.Old,
		"new": aa.Property.New,
	}

	activity.CausedBy = data.Employee.ID
	activity.CausedByName = data.Employee.FullName

	err := Config.PgSQL.Create(&activity).Error
	if err != nil {
		xtremelog.Error(err)
		Error.ErrActivitySave(err.Error())
	}

	return nil
}

func (aa UseActivity) setPropertyWithParser(action string, subs ...string) interface{} {
	var property interface{}
	var subAction string

	if len(subs) > 0 {
		subAction = subs[0]
	}

	var parser ModelActivityProperty
	parser.SetParser(aa.Model)

	switch action {
	case Server.ACTION_CREATE:
		property = parser.Base.CreateActivity(subAction)
		break
	case Server.ACTION_UPDATE:
		property = parser.Base.UpdateActivity(subAction)
		break
	case Server.ACTION_DELETE:
		property = parser.Base.DeleteActivity(subAction)
		break
	case Server.ACTION_GENERAL:
		property = parser.Base.GeneralActivity(subAction)
		break
	default:
		Error.ErrActivityActionType()
		break
	}

	return property
}
