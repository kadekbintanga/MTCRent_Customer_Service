package model

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
)

type Customer struct {
	xtrememodel.BaseModelUUID
	Name            string                          `gorm:"column:name;type:varchar(250);not null"`
	IDNumber        string                          `gorm:"column:IDNumber;type:varchar(100);not null"`
	SIMNumber       string                          `gorm:"column:SIMNumber;type:varchar(100);not null"`
	Phone           string                          `gorm:"column:phone;type:varchar(30);not null"`
	Address         string                          `gorm:"column:address;type:varchar(250);default:null"`
	StatusId        int                             `gorm:"column:statusId"`
	BlacklistReason *string                         `gorm:"column:blacklistReason;default:null"`
	IdentityPhoto   *xtrememodel.MapInterfaceColumn `gorm:"column:identityPhoto;type:json;null"`
	CreatedBy       *string                         `gorm:"column:createdBy;varchar(50);null"`
	CreatedByName   *string                         `gorm:"column:createdByName;varchar(250);null"`
	UpdatedBy       *string                         `gorm:"column:updatedBy;varchar(50);null"`
	UpdatedByName   *string                         `gorm:"column:updatedByName;varchar(250);null"`
}

func (Customer) TableName() string {
	return "customers"
}

func (model Customer) SetReference() uint {
	return model.ID
}
