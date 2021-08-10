package token

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/common/model/dto/user"
	errConsts "github.com/ravielze/oculi/constant/errors"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/context"
	errorUtil "github.com/ravielze/oculi/errors"
	"github.com/ravielze/oculi/response"
	"github.com/ravielze/oculi/token"
)

type (
	ClaimFunction func(req *http.Request) (token.Claims, error)
)

var (
	tokenMiddlewareActivated = false
)

func EchoMiddleware(cf ClaimFunction) echo.MiddlewareFunc {
	tokenMiddlewareActivated = true
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*context.Context)
			if !ok {
				ctx = context.New(c)
			}
			claims, err := cf(ctx.Request())
			if err != nil {
				ctx.Set(consts.KeyCredentials, user.CredentialsDTO{
					ID:       0,
					Metadata: errorUtil.InjectDetails(errConsts.ErrUnauthorized, err.Error()),
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
				ctx.AddError(
					http.StatusUnauthorized,
					errorUtil.InjectDetails(
						errConsts.ErrUnauthorized,
						"unauthorized: credentials not found",
					),
				)
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
