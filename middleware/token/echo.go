package token

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/token"
)

var (
	tokenMiddlewareActivated = false
)

func EchoMiddleware(token token.Tokenizer) echo.MiddlewareFunc {
	tokenMiddlewareActivated = true
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*context.Context)
			if !ok {
				ctx = context.New(c)
			}
			claims, err := token.DecodeHttpRequest(ctx.Request())
			if err != nil {
				ctx.AddError(http.StatusUnauthorized, fmt.Errorf("unauthorized: %s", err.Error()))
				return next(c)
			}
			ctx.Set(consts.KeyCredentials, claims.Credentials())
			return next(c)
		}
	}
}
