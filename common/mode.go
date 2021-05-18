package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func DevMode() bool {
	godotenv.Load()
	mode := os.Getenv("MODE")
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
