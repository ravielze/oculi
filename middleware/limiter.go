package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common/code"
	"github.com/ravielze/oculi/common/utils"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func InstallLimiter(g *gin.Engine, r rate.Limit, burst int, dur time.Duration) {
	g.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(r, burst), dur
	}, func(c *gin.Context) {
		utils.AbortAndResponse(c, http.StatusTooManyRequests, code.TOO_MANY_REQUESTS)
	}))
}

func InstallDefaultLimiter(g *gin.Engine) {
	g.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(250*time.Millisecond), 5), time.Minute * 15
	}, func(c *gin.Context) {
		utils.AbortAndResponse(c, http.StatusTooManyRequests, code.TOO_MANY_REQUESTS)
	}))
}
