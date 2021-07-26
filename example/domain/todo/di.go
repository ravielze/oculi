package todo

import (
	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/domain/todo/repository"
	"github.com/ravielze/oculi/example/domain/todo/service"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Provide(repository.New).
		Provide(service.New).
		End()
}
