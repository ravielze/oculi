package database

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySql(config *gorm.Config) *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	timeZone := os.Getenv("DB_TIMEZONE")
	timeZone = strings.ReplaceAll(timeZone, "/", "%2F")
	loginInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&loc=%s", user, password, host, port, dbname, timeZone)
	db, err := gorm.Open(mysql.Open(loginInfo), config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}
