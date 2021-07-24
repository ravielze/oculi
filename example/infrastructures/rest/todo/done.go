package todo

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Done(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database).Param("id")

	result := ctx.Process(c.Handler.Todo.Done(req))

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
