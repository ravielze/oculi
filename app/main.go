package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ravielze/fuzzy-broccoli/module"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	conn "github.com/ravielze/fuzzy-broccoli/connection"
)

func main() {
	godotenv.Load()

	serverMode := GetServerMode(os.Getenv("MODE"))

	db := conn.ConnectDatabase(os.Getenv("DB_DRIVER"), serverMode)
	engine := gin.Default()
	modules := module.NewModule(db, engine)
	fmt.Println("\033[33mModule Installed:\033[0m")
	for i, x := range modules {
		fmt.Printf("\033[34m[%d] \033[0m: \033[32m%s\033[0m\n", i, x.Name())
	}
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
	} else {
		fmt.Println("Starting production server...")
		gin.SetMode(gin.ReleaseMode)
	}
	return result
}
