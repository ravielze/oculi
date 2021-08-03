package user

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/example/constants"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Check(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database).Transform()

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.User.Check, req),
		nil,
		constants.UserMappers,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, nil, result)
}
