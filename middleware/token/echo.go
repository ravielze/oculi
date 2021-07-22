package token

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/response"
	"github.com/ravielze/oculi/token"
)

const (
	CredentialsKey = "CREDENTIALS"
)

func EchoMiddleware(token token.Tokenizer, response response.Responder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*context.Context)
			claims, err := token.DecodeHttpRequest(ctx.Request())
			if err != nil {
				ctx.AddError(http.StatusUnauthorized, fmt.Errorf("unauthorized: %s", err.Error()))
				return response.NewJSONResponse(ctx, nil, nil)
			}
			ctx.Set(CredentialsKey, claims.Credentials())
			return next(c)
		}
	}
}
