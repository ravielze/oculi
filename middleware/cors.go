package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InstallCors(g *gin.Engine) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           15 * time.Minute,
	}
	g.Use(cors.New(corsConfig))
}
