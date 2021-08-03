package handlers

import (
	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/handlers/health"
	"github.com/ravielze/oculi/example/handlers/todo"
	"github.com/ravielze/oculi/example/handlers/user"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	return di.NewRegistrant(container).
		Provide(health.NewHandler).
		Provide(user.NewHandler).
		Provide(todo.NewHandler).
		Proceed()
}
