package Testing

import (
	"Service/App/Model"
)

type TestingSub struct {
	Model.BaseModel
	TestingId uint   `gorm:"column:testingId;type:bigint"`
	Name      string `gorm:"column:name;type:varchar(250);default:null"`
}

func (TestingSub) TableName() string {
	return "testing_subs"
}
