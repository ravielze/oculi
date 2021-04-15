package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
)


func AbortAndResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, serializer.NewResponse(code, msg))
}

func AbortAndResponseData(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.AbortWithStatusJSON(code, serializer.NewResponseData(code, msg, data))
}

func OKAndResponse(ctx *gin.Context){
	ctx.JSON(http.StatusOK, serializer.NewResponse(http.StatusOK, code.OK))
}

func OKAndResponseData(ctx *gin.Context, data interface{}){
	ctx.JSON(http.StatusOK, serializer.NewResponseData(http.StatusOK, code.OK, data))
}