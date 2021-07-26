package service

import (
	"github.com/ravielze/oculi/example/domain/user/repository"
	"github.com/ravielze/oculi/example/model/dao"
	userDto "github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/hash"
	"github.com/ravielze/oculi/request"
)

type (
	Service interface {
		Login(req request.Context, item userDto.LoginRequest) (user dao.User, token string, err error)
		Register(req request.Context, user userDto.RegisterRequest) error
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
		hash       hash.Hash
	}
)

func New(r resources.Resource, repo repository.Repository, h hash.Hash) Service {
	return &service{
		resource:   r,
		repository: repo,
		hash:       h,
	}
}
