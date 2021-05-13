package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InstallCors(g *gin.Engine, origins []string) {

	corsConfig := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           15 * time.Minute,
		//AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
	}
	// g.Use(func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	// })
	g.Use(cors.New(corsConfig))
}
