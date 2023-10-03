package Testing

import (
	"Service/App/Models"
)

type Testing struct {
	Models.BaseModel
	Name string `gorm:"column:name;type:varchar(250);default:null"`
}

func (Testing) TableName() string {
	return "testing"
}

func (model Testing) SetProperty() map[string]interface{} {
	return map[string]interface{}{
		"name": model.Name,
	}
}

func (model Testing) SetReference() uint {
	return model.BaseModel.ID.ID
}
