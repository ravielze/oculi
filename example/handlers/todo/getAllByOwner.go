package todo

import (
	todoDto "github.com/ravielze/oculi/example/model/dto/todo"
	"github.com/ravielze/oculi/request"
)

func (h *handler) GetAllByOwner(req request.ReqContext) (todoDto.TodosResponse, error) {
	todos, err := h.domain.Todo.GetAllByOwner(req)
	if err != nil {
		return todoDto.TodosResponse{}, err
	}
	return todoDto.NewTodosResponse(todos), nil
}
