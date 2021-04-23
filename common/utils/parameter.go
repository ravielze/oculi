package utils

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/radix36"
)

func GetParam(ctx *gin.Context, parameter string) (string, bool) {
	result := ctx.Param(parameter)
	if len(result) == 0 {
		AbortAndResponseData(
			ctx,
			http.StatusUnprocessableEntity,
			code.PARAMETER_ERROR,
			fmt.Sprintf("parameter '%s' is missing", parameter),
		)
		return "", false
	}
	return result, true
}

func GetParamID(ctx *gin.Context, parameter string) (uuid.UUID, bool) {
	parameterGetter := ctx.Param(parameter)
	if len(parameterGetter) == 0 {
		AbortAndResponseData(
			ctx,
			http.StatusUnprocessableEntity,
			code.PARAMETER_ERROR,
			fmt.Sprintf("parameter '%s' is missing", parameter),
		)
		return uuid.Nil, false
	}
	result := radix36.DecodeUUID(parameterGetter)
	if result != uuid.Nil {
		AbortAndResponseData(
			ctx,
			http.StatusUnprocessableEntity,
			code.PARAMETER_ERROR,
			fmt.Sprintf("parameter '%s' is not radix36", parameter),
		)
		return uuid.Nil, false
	}
	return result, true
}
