package common

import (
	"strings"
	"time"

	"github.com/ravielze/fuzzy-broccoli/common/radix36"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IntIDBase struct {
	ID uint `gorm:"type:BIGINT;primaryKey;autoIncrement;uniqueIndex:,sort:asc,type:btree"`
}

type BigIntIDBase struct {
	ID uint64 `gorm:"type:BIGINT;primaryKey;autoIncrement;uniqueIndex:,sort:asc,type:btree"`
}

type StringIDBase struct {
	ID string `gorm:"primaryKey;type:VARCHAR(25);uniqueIndex:,sort:asc,type:btree"`
}

type UUIDBase struct {
	ID string `gorm:"primaryKey;type:VARCHAR(32);uniqueIndex:,sort:asc,type:btree"`
}

func (e *UUIDBase) BeforeCreate(scope *gorm.DB) error {
	if strings.Contains(e.ID, "default") {
		return nil
	}
	uuid := uuid.NewV4().String()
	e.ID = strings.ToUpper(strings.Replace(uuid, "-", "", 4))
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
