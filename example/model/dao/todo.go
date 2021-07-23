package dao

import (
	"github.com/ravielze/oculi/common/model/dao"
)

type (
	Todo struct {
		ID                      uint64  `gorm:"primaryKey"`
		OwnerID                 uint64  `gorm:"not null"`
		Title                   string  `gorm:"not null;type:VARCHAR(256)"`
		Description             *string `gorm:"null;type:VARCHAR(1024)"`
		IsDone                  bool    `gorm:"not null"`
		dao.BaseModelSoftDelete `gorm:"embedded"`

		Owner User `gorm:"->;foreignKey:OwnerID;references:ID"`
	}
)
