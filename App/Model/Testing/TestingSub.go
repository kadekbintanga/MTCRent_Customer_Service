package Testing

import (
	"github.com/globalxtreme/gobaseconf/model"
)

type TestingSub struct {
	model.BaseModel
	TestingId uint   `gorm:"column:testingId;type:bigint"`
	Name      string `gorm:"column:name;type:varchar(250);default:null"`
}

func (TestingSub) TableName() string {
	return "testing_subs"
}
