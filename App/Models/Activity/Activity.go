package Activity

import (
	"Service/App/Models"
)

type Activity struct {
	Models.BaseModel
	Feature      string           `gorm:"column:feature;type:varchar(150)"`
	SubFeature   string           `gorm:"column:subFeature;type:varchar(150);default:null"`
	Action       string           `gorm:"column:action;type:varchar(50)"`
	Description  string           `gorm:"column:description;type:text;default:null"`
	Reference    uint             `gorm:"column:reference;type:integer;default:null"`
	CausedBy     string           `gorm:"column:causedBy;type:varchar(50);default:null"`
	CausedByName string           `gorm:"column:causedByName;type:varchar(150);default:null"`
	Properties   Models.MapColumn `gorm:"column:properties;type:json;default:null"`
}

func (Activity) TableName() string {
	return "activities"
}
