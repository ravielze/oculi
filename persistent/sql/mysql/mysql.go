package mysql

import (
	"github.com/ravielze/oculi/persistent/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(connInfo sql.ConnectionInfo) gorm.Dialector {
	return mysql.Open(connInfo.URI())
}
