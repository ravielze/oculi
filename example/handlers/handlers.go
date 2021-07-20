package handlers

import (
	"github.com/ravielze/oculi/example/handlers/health"
	"go.uber.org/dig"
)

type (
	Handler struct {
		dig.In

		Health health.Handler
	}
)
