package sqlv2

import (
	"github.com/ravielze/oculi/persistent/sql"
	"gorm.io/gorm"
)

func NewClient(dl gorm.Dialector, config *gorm.Config) (sql.API, error) {
	db, err := gorm.Open(dl, config)
	if err != nil {
		return nil, err
	}

	return &Impl{
		Database: db,
	}, nil
}
