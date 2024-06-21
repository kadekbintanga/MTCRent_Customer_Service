package activity

import (
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/model"
)

type ActivityModelInterface interface {
	TableName() string
	SetReference() string
}

type property struct {
	Old interface{}
	New interface{}
}

type UseActivity struct {
	Reference   string   `gorm:"-"`
	Feature     string   `gorm:"-"`
	SubFeature  string   `gorm:"-"`
	Action      string   `gorm:"-"`
	Description string   `gorm:"-"`
	Property    property `gorm:"-"`
	Parser      core.BaseActivityPropertyParserInterface
	//Employee    data.EmployeeIdentifierData    `gorm:"-"` // TODO: Re-enable this code after installing github.com/globalxtreme/go-identifier module (If you use GX Identifier for authorization)
}

func (aa UseActivity) SetReference(md ActivityModelInterface) UseActivity {
	aa.Feature = md.TableName()
	aa.Reference = md.SetReference()

	return aa
}

func (aa UseActivity) SetSubFeature(subFeature string) UseActivity {
	aa.SubFeature = subFeature

	return aa
}

func (aa UseActivity) SetParser(Parser core.BaseActivityPropertyParserInterface) UseActivity {
	aa.Parser = Parser

	return aa
}

func (aa UseActivity) SetOldProperty(action string, subs ...string) UseActivity {
	aa.Action = action

	if aa.Parser != nil {
		aa.Property.Old = aa.setPropertyWithParser(action, subs...)
	}

	return aa
}

func (aa UseActivity) SetNewProperty(action string, subs ...string) UseActivity {
	aa.Action = action

	if aa.Parser != nil {
		aa.Property.New = aa.setPropertyWithParser(action, subs...)
	}

	return aa
}

func (aa UseActivity) Save(description string) error {
	var activity model.Activity
	activity.Feature = aa.Feature
	activity.SubFeature = aa.SubFeature
	activity.Action = aa.Action
	activity.Description = description
	activity.Reference = aa.Reference
	activity.Properties = map[string]interface{}{
		"old": aa.Property.Old,
		"new": aa.Property.New,
	}

	// TODO: Re-enable this code after installing github.com/globalxtreme/go-identifier module (If you use GX Identifier for authorization)
	//activity.CausedBy = aa.Employee.ID
	//activity.CausedByName = aa.Employee.FullName

	err := config.PgSQL.Create(&activity).Error
	if err != nil {
		xtremelog.Error(err)
		error2.ErrXtremeActivitySave(err.Error())
	}

	return nil
}

func (aa UseActivity) setPropertyWithParser(action string, subs ...string) interface{} {
	var property interface{}
	var subAction string

	if len(subs) > 0 {
		subAction = subs[0]
	}

	switch action {
	case constant.ACTION_CREATE:
		property = aa.Parser.CreateActivity(subAction)
	case constant.ACTION_UPDATE:
		property = aa.Parser.UpdateActivity(subAction)
	case constant.ACTION_DELETE:
		property = aa.Parser.DeleteActivity(subAction)
	case constant.ACTION_GENERAL:
		property = aa.Parser.GeneralActivity(subAction)
	default:
		error2.ErrXtremeActivityActionType()
	}

	return property
}
