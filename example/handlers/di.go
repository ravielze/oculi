package handlers

import (
	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/handlers/health"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	return di.NewRegistrant(container).
		Provide(health.NewHandler).
		End()
}
