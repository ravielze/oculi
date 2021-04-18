package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
)

var token2 string

func resetToken(ctx *gin.Context) {
	if values := ctx.Request.Header.Get("Authorization"); len(values) > 0 {
		if values == token2 {
			ctx.Next()
			return
		}

	}

	ctx.AbortWithStatusJSON(http.StatusForbidden, serializer.NewResponse(http.StatusForbidden, code.UNAUTHORIZED))
}

func GetResetTokenMiddleware() gin.HandlerFunc {
	token2 = os.Getenv("RESET_TOKEN")
	return resetToken
}