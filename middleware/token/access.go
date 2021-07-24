package token

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/response"
)

func PublicEndpoint() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*context.Context)
			if len(ctx.Errors()) == 1 && ctx.ResponseCode() == http.StatusUnauthorized {
				ctx.ClearErrors()
			}
			return next(c)
		}
	}
}

func PrivateEndpoint(r response.Responder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(*context.Context)
			if ctx.HasError() && ctx.ResponseCode() == http.StatusUnauthorized {
				return r.NewJSONResponse(ctx, nil, nil)
			}
			return next(c)
		}
	}
}
