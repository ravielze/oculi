package user

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	dto "github.com/ravielze/oculi/example/model/dto/user"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Login(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.LoginRequest
	ctx.BindValidate(&item)

	result := ctx.Process(c.Handler.User.Login(req, item))

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
