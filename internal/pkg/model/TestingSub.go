package model

import (
	base "github.com/globalxtreme/gobaseconf/model"
)

type TestingSub struct {
	base.BaseModel
	TestingId string `gorm:"column:testingId;type:varchar(45);not null"`
	Name      string `gorm:"column:name;type:varchar(250);default:null"`
}

func (TestingSub) TableName() string {
	return "testing_subs"
}
