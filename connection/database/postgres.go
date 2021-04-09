package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres(config *gorm.Config) *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	timeZone := os.Getenv("DB_TIMEZONE")
	loginInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, dbname, timeZone)

	db, err := gorm.Open(postgres.Open(loginInfo), config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}
