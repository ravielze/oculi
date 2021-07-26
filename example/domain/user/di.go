package user

import (
	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/domain/user/repository"
	"github.com/ravielze/oculi/example/domain/user/service"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Provide(repository.New).
		Provide(service.New).
		End()
}
