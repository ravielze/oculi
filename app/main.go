package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	conn "github.com/ravielze/fuzzy-broccoli/connection"
	"github.com/ravielze/fuzzy-broccoli/module/auth"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	serverMode := GetServerMode(os.Getenv("MODE"))

	var db *gorm.DB
	conn.ConnectDatabase(os.Getenv("DB_DRIVER"), db, serverMode)
	engine := gin.Default()
	auth.NewAuthModule(db, engine)
	engine.Run()
}

func GetServerMode(mode string) bool {
	result := true
	switch {
	case strings.EqualFold(mode, "production"):
		result = false
	case strings.EqualFold(mode, "prod"):
		result = false
	case strings.EqualFold(mode, "p"):
		result = false
	case strings.EqualFold(mode, "1"):
		result = false
	case strings.EqualFold(mode, "development"):
	case strings.EqualFold(mode, "dev"):
	case strings.EqualFold(mode, "d"):
	case strings.EqualFold(mode, "0"):
	default:
		panic("MODE only can be development/production")
	}
	if result {
		fmt.Println("Starting development server...")
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("Starting production server...")
	}
	return result
}
