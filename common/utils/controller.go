package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	code "github.com/ravielze/oculi/common/code"
	"github.com/ravielze/oculi/common/serializer"
)

func AbortAndResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, serializer.NewResponse(code, msg))
}

func AbortAndResponseData(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.AbortWithStatusJSON(code, serializer.NewResponseData(code, msg, data))
}

func AbortUsecaseError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest,
		serializer.NewResponseData(http.StatusBadRequest,
			code.LOGIC_ERROR, err.Error(),
		),
	)
}

func OKAndResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		serializer.NewResponse(http.StatusOK, code.OK),
	)
}

func OKAndResponseData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK,
		serializer.NewResponseData(http.StatusOK, code.OK, data),
	)
}