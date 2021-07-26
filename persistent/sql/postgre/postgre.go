package postgre

import (
	"github.com/ravielze/oculi/persistent/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(connInfo sql.ConnectionInfo) gorm.Dialector {
	return postgres.Open(connInfo.PostgresURI())
}
