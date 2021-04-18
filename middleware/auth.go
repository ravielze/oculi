package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	"github.com/ravielze/fuzzy-broccoli/module/auth"
)

var AuthModule *auth.AuthModule

func GetAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := auth.ExtractTokenID(ctx.Request)
		if err != nil {
			utils.AbortAndResponse(ctx, http.StatusUnauthorized, code.UNAUTHORIZED)
			return
		}
		user, err2 := (*AuthModule).Usecase.GetID(userId)
		if err2 != nil {
			utils.AbortAndResponseData(ctx, http.StatusUnauthorized, code.UNAUTHORIZED, err2.Error())
			return
		}
		ctx.Keys = map[string]interface{}{}
		ctx.Keys["user"] = user
	}
}

func GetUser(ctx *gin.Context) auth.User {
	if user, ok := ctx.Keys["user"].(auth.User); ok {
		return user
	}
	return auth.User{}
}
