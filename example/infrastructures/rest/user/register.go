package user

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/example/constants"
	dto "github.com/ravielze/oculi/example/model/dto/user"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) RegisterUser(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.RegisterRequest
	ctx.BindValidate(&item)

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.User.Register, req, item),
		nil,
		constants.UserMappers,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
