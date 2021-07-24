package todo

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	dto "github.com/ravielze/oculi/example/model/dto/todo"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Create(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.CreateTodoRequest
	ctx.BindValidate(&item)

	result := ctx.Process(c.Handler.Todo.Create(req, item))

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
