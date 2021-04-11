package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
)

var token string

func StaticToken(ctx *gin.Context) {
	if values := ctx.Request.Header.Get("Authorization"); len(values) > 0 {
		if values == token {
			ctx.Next()
			return
		}

	}

	ctx.AbortWithStatusJSON(http.StatusForbidden, serializer.NewResponse(http.StatusForbidden, code.UNAUTHORIZED))
}

func InstallStaticToken(rg *gin.RouterGroup) {
	token = os.Getenv("AUTH_TOKEN")
	rg.Use(StaticToken)
}

func InstallGlobalStaticToken(g *gin.Engine) {
	token = os.Getenv("AUTH_TOKEN")
	g.Use(StaticToken)
}
