package user

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Check(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	result := ctx.Process(c.Handler.User.Check(req))

	return c.Resource.Responder.NewJSONResponse(ctx, nil, result)
}
