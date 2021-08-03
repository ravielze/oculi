package repository

import (
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		Create(req request.ReqContext, item dao.Todo) (dao.Todo, error)
		GetByID(req request.ReqContext, todoId uint64) (dao.Todo, error)
		Delete(req request.ReqContext, todoId uint64) error
		Update(req request.ReqContext, todoId uint64, request dto.Map) error
		GetAllByOwner(req request.ReqContext) ([]dao.Todo, error)
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
