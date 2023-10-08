package Model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/* --- BASE MODEL CONFIGURATION --- */

type BaseModel struct {
	ID
	DateTime
}

type ID struct {
	ID uint `gorm:"primarykey"`
}

type DateTime struct {
	CreatedAt time.Time      `gorm:"column:createdAt;type:timestamp"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;index"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

/* --- COLUMN TYPE CONFIGURATION: OBJECT / MAP IN ARRAY --- */

type ArrayMapColumn []map[string]interface{}

func (j *ArrayMapColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result []map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j ArrayMapColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}

type MapColumn map[string]interface{}

func (j *MapColumn) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var result map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

func (j MapColumn) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.Marshal(j)
}
