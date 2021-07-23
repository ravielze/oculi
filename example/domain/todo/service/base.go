package service

import (
	"github.com/ravielze/oculi/example/domain/todo/repository"
	"github.com/ravielze/oculi/example/resources"
)

type (
	Service interface {
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
