package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common/code"
	"github.com/ravielze/oculi/common/serializer"
)

func GetStaticTokenMiddleware() gin.HandlerFunc {
	token := os.Getenv("AUTH_TOKEN")
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
