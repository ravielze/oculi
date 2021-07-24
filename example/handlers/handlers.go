package handlers

import (
	"github.com/ravielze/oculi/example/handlers/health"
	"github.com/ravielze/oculi/example/handlers/user"
	"go.uber.org/dig"
)

type (
	Handler struct {
		dig.In

		Health health.Handler
		User   user.Handler
	}
)
