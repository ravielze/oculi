package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	BaseInfo
}

func (dbdriver *PostgreSQL) Connect(config *gorm.Config) *gorm.DB {
	loginInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		dbdriver.Host, dbdriver.Port, dbdriver.User, dbdriver.Password, dbdriver.Database, dbdriver.TimeZone)
	db, err := gorm.Open(postgres.Open(loginInfo), config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}
