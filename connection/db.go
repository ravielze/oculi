package conn

import (
	"strings"

	"github.com/ravielze/fuzzy-broccoli/connection/database"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(dbdriver string, development bool) *gorm.DB {
	config := &gorm.Config{
		// Enable this if there is any nullable foreign key
		// Don't forget to set up relation manually
		// DisableForeignKeyConstraintWhenMigrating: true,
	}
	if development {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	var db *gorm.DB
	switch {
	case strings.EqualFold(dbdriver, "mysql"):
		db = database.MySql(config)
	case strings.EqualFold(dbdriver, "postgres"):
		db = database.Postgres(config)
	default:
		panic("Database Driver is not found.")
	}
	if db == nil {
		panic("Database is not connected.")
	}
	return db
}
