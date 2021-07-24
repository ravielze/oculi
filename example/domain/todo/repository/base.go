package repository

import (
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		Create(req request.Context, item dao.Todo) (dao.Todo, error)
		GetByID(req request.Context, todoId uint64) (dao.Todo, error)
		Delete(req request.Context, todoId uint64) error
		Update(req request.Context, todoId uint64, request dto.Map) error
		GetAllByOwner(req request.Context) ([]dao.Todo, error)
	}

	repository struct {
		resource resources.Resource
	}
)

func New(r resources.Resource) Repository {
	return &repository{
		resource: r,
	}
}
