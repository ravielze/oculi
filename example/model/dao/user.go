package dao

import "github.com/ravielze/oculi/common/model/dao"

type (
	User struct {
		ID            uint64 `gorm:"primaryKey" json:"id"`
		Username      string `gorm:"not null;uniqueIndex:username;type:VARCHAR(128);" json:"username"`
		Password      string `gorm:"not null;type:VARCHAR(128);" json:"-"`
		dao.BaseModel `gorm:"embedded"`
	}
)
