package database

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	BaseInfo
}

func (dbdriver *MySQL) Connect(config *gorm.Config) *gorm.DB {
	dbdriver.TimeZone = strings.ReplaceAll(dbdriver.TimeZone, "/", "%2F")
	loginInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&loc=%s",
		dbdriver.User, dbdriver.Password, dbdriver.Host, dbdriver.Port, dbdriver.Database, dbdriver.TimeZone)
	db, err := gorm.Open(mysql.Open(loginInfo), config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}
