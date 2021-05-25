package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InstallCors(g *gin.Engine) {

	g.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	})
	corsConfig := cors.Config{
		AllowMethods:           []string{"PUT", "GET", "POST", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Accept"},
		AllowCredentials:       true,
		MaxAge:                 1 * time.Minute,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowAllOrigins:        true,
	}
	g.Use(cors.New(corsConfig))
}
