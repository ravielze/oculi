package todo

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/example/constants"
	dto "github.com/ravielze/oculi/example/model/dto/todo"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Edit(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.UpdateTodoRequest
	ctx.BindValidate(&item)

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Todo.Edit, req, item),
		nil,
		constants.TodoMappers,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
