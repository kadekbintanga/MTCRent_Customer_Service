package model

import (
	base "github.com/globalxtreme/gobaseconf/model"
)

type Testing struct {
	base.BaseModel
	Name string       `gorm:"column:name;type:varchar(250);default:null"`
	Subs []TestingSub `gorm:"foreignKey:testingId"`
}

func (Testing) TableName() string {
	return "testing"
}

func (model Testing) SetReference() string {
	return model.BaseModel.ID
}
