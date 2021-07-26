package dao

import (
	"time"

	"gorm.io/gorm"
)

type (
	BaseModel struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	BaseModelSoftDelete struct {
		BaseModel
		DeletedAt gorm.DeletedAt
	}

	Map map[string]interface{}
)

func (m Map) ToMap() map[string]interface{} {
	return m
}
