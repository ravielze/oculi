package sqlv2

import (
	"gorm.io/gorm"
)

type (
	Impl struct {
		Database *gorm.DB
		Object   []interface{}
	}
)

func (i *Impl) copy(db *gorm.DB) *Impl {
	return &Impl{
		Database: db,
		Object:   i.Object,
	}
}
