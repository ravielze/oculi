package common

import (
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/radix36"
	"gorm.io/gorm"
)

type IntIDBase struct {
	ID uint `gorm:"primaryKey;autoIncrement;uniqueIndex:,sort:asc,type:btree"`
}

type BigIntIDBase struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;uniqueIndex:,sort:asc,type:btree"`
}

type StringIDBase struct {
	ID string `gorm:"primaryKey;type:VARCHAR(25);uniqueIndex:,sort:asc,type:btree"`
}

type UUIDBase struct {
	ID string `gorm:"primaryKey;type:VARCHAR(36);uniqueIndex:,sort:asc,type:btree"`
}

func (e *UUIDBase) BeforeCreate(scope *gorm.DB) error {
	if strings.Contains(e.ID, "default") {
		return nil
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	e.ID = uuid.String()
	return nil
}

func (e *StringIDBase) BeforeCreate(scope *gorm.DB) error {
	if strings.EqualFold(e.ID, "default") {
		return nil
	}
	id, err := radix36.EncodeUUID4()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

type InfoBase struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDeleteBase struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
