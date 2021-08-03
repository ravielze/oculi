package service

import (
	"github.com/ravielze/oculi/example/domain/todo/repository"
	"github.com/ravielze/oculi/example/model/dao"
	todoDto "github.com/ravielze/oculi/example/model/dto/todo"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Service interface {
		Create(req request.ReqContext, todo todoDto.CreateTodoRequest) (dao.Todo, error)
		Done(req request.ReqContext, todoId uint64) error
		Undone(req request.ReqContext, todoId uint64) error
		Edit(req request.ReqContext, todo todoDto.UpdateTodoRequest) error
		Delete(req request.ReqContext, todoId uint64) error
		GetAllByOwner(req request.ReqContext) ([]dao.Todo, error)
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
	}
)

func New(r resources.Resource, repo repository.Repository) Service {
	return &service{
		resource:   r,
		repository: repo,
	}
}
