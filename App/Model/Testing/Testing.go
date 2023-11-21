package Testing

import (
	"github.com/globalxtreme/gobaseconf/model"
)

type Testing struct {
	model.BaseModel
	Name string       `gorm:"column:name;type:varchar(250);default:null"`
	Subs []TestingSub `gorm:"foreignKey:testingId"`
}

func (Testing) TableName() string {
	return "testing"
}

func (model Testing) SetReference() uint {
	return model.BaseModel.ID
}
