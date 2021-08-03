package repository

import (
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		Create(req request.ReqContext, user dao.User) (dao.User, error)
		GetByUsername(req request.ReqContext, username string) (dao.User, error)
		GetByID(req request.ReqContext, userId uint64) (dao.User, error)
		Update(req request.ReqContext, userId uint64, request dto.Map) error
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
