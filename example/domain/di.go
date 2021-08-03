package domain

import (
	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/domain/todo"
	"github.com/ravielze/oculi/example/domain/user"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Register(user.Register).
		Register(todo.Register).
		Proceed()
}
