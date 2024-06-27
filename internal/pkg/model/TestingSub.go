package model

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type TestingSub struct {
	xtrememodel.BaseModel
	TestingId string `gorm:"column:testingId;type:varchar(45);not null"`
	Name      string `gorm:"column:name;type:varchar(250);default:null"`
}

func (TestingSub) TableName() string {
	return "testing_subs"
}
