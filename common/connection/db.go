package conn

import (
	"os"
	"strings"

	"github.com/ravielze/oculi/common/connection/database"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConnector interface {
	Connect(config *gorm.Config) *gorm.DB
}

func ConnectDatabase(development bool) *gorm.DB {
	databaseDriver := os.Getenv("DB_DRIVER")
	config := &gorm.Config{
		// Enable this if there is any nullable foreign key
		// Don't forget to set up relation manually
		// DisableForeignKeyConstraintWhenMigrating: true,
	}
	if development {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	databaseInfo := database.BaseInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	var databaseConnector DatabaseConnector

	var db *gorm.DB
	switch {
	case strings.EqualFold(databaseDriver, "mysql"):
		databaseConnector = &database.MySQL{
			BaseInfo: databaseInfo,
		}
	case strings.EqualFold(databaseDriver, "postgres"):
		databaseConnector = &database.PostgreSQL{
			BaseInfo: databaseInfo,
		}
	default:
		panic("Database Driver is not found.")
	}

	db = databaseConnector.Connect(config)
	if db == nil {
		panic("Database is not connected.")
	}
	return db
}
