package sqlv2

import (
	"gorm.io/gorm"
)

type (
	Impl struct {
		Database *gorm.DB
	}
)

func (i *Impl) copy(db *gorm.DB) *Impl {
	return &Impl{
		Database: db,
	}
}
