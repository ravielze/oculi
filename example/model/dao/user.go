package dao

import "github.com/ravielze/oculi/common/model/dao"

type (
	User struct {
		ID            uint64 `gorm:"primaryKey"`
		Username      string `gorm:"not null;uniqueIndex:username;type:VARCHAR(128);"`
		Password      string `gorm:"not null;type:VARCHAR(128);"`
		dao.BaseModel `gorm:"embedded"`
	}
)
