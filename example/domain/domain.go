package domain

import (
	todoService "github.com/ravielze/oculi/example/domain/todo/service"
	userService "github.com/ravielze/oculi/example/domain/user/service"
	"go.uber.org/dig"
)

type (
	Domain struct {
		dig.In

		User userService.Service
		Todo todoService.Service
	}
)
