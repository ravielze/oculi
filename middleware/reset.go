package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	"gorm.io/gorm"
)

func GetResetTokenMiddleware() gin.HandlerFunc {
	token := os.Getenv("RESET_TOKEN")
	return func(ctx *gin.Context) {
		if values := ctx.Request.Header.Get("Authorization"); len(values) > 0 {
			if values == token {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, serializer.NewResponse(http.StatusForbidden, code.UNAUTHORIZED))
	}
}

func GetResetHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.OKAndResponseData(ctx, "Server will be reseted in 3 seconds")
		var channel = make(chan bool)
		go delayReset(channel)
		go checkForReset(channel, db)
	}
}

func delayReset(ch chan<- bool) {
	time.Sleep(3 * time.Second)
	ch <- true
}

type ResetHandler func(db *gorm.DB)

var ResetFunction ResetHandler

func checkForReset(ch <-chan bool, db *gorm.DB) {
loop:
	for {
		data := <-ch
		if data {
			ResetFunction(db)
			break loop
		}
	}
}
