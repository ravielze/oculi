package token

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/common/model/dto/user"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/response"
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
				ctx.Set(consts.KeyCredentials, user.CredentialsDTO{
					ID:       0,
					Metadata: fmt.Errorf("unauthorized: %s", err.Error()),
				})
			} else {
				ctx.Set(consts.KeyCredentials, claims.Credentials())
			}
			return next(c)
		}
	}
}

func PrivateEndpoint(r response.Responder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		if !tokenMiddlewareActivated {
			return func(c echo.Context) error {
				return next(c)
			}
		}
		return func(c echo.Context) error {
			ctx := c.(*context.Context)
			credentials := ctx.Get(consts.KeyCredentials)
			if credentials == nil {
				ctx.AddError(http.StatusUnauthorized, errors.New("unauthorized: credentials not found"))
				return r.NewJSONResponse(ctx, nil, nil)
			} else {
				cdto := credentials.(user.CredentialsDTO)
				if cdto.ID == 0 {
					ctx.AddError(http.StatusUnauthorized, cdto.Metadata.(error))
					return r.NewJSONResponse(ctx, nil, nil)
				}
			}
			return next(c)
		}
	}
}
