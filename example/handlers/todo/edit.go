package todo

import (
	todoDto "github.com/ravielze/oculi/example/model/dto/todo"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Edit(req request.ReqContext, item todoDto.UpdateTodoRequest) error {
	return h.domain.Todo.Edit(req, item)
}
