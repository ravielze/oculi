package todo

import (
	"github.com/ravielze/oculi/example/domain"
	todoDto "github.com/ravielze/oculi/example/model/dto/todo"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	handler struct {
		domain   domain.Domain
		resource resources.Resource
	}

	Handler interface {
		Create(req request.Context, item todoDto.CreateTodoRequest) (todoDto.TodoResponse, error)
		Done(req request.Context) error
		Undone(req request.Context) error
		Edit(req request.Context, item todoDto.UpdateTodoRequest) error
		Delete(req request.Context) error
		GetAllByOwner(req request.Context) (todoDto.TodosResponse, error)
	}
)

func NewHandler(domain domain.Domain, resource resources.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
