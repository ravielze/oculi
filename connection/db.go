package conn

import (
	"strings"

	"gorm.io/gorm"
)

func ConnectDatabase(dbms string, db *gorm.DB) {
	switch {
	case strings.EqualFold(dbms, "mysql"):
		//TODO create connector
	case strings.EqualFold(dbms, "postgres"):
		//TODO create connector
	default:
		panic("DBMS connector not found.")
	}
	if db == nil {
		panic("Database is not connected.")
	}
}
