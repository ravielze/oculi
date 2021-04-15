package common

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IDBase struct {
	ID uint `gorm:"primaryKey;autoIncrement;uniqueIndex:,sort:asc,type:btree" json:"id"`
}

type UUIDBase struct {
	ID string `gorm:"primaryKey;type:VARCHAR(36);uniqueIndex:,sort:asc,type:btree" json:"id"`
}

func (e *UUIDBase) BeforeCreate(scope *gorm.DB) error {
	e.ID = uuid.NewV4().String()
	return nil
}

type InfoBase struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDeleteBase struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}