package rest

import (
	"github.com/ravielze/oculi/example/infrastructures/rest/health"
	"github.com/ravielze/oculi/example/infrastructures/rest/user"
	"github.com/ravielze/oculi/example/resources"
	"go.uber.org/dig"
)

type (
	Rest struct {
		dig.In

		Controller Controller
		Resource   resources.Resource
	}

	Controller struct {
		dig.In

		Health health.Controller
		User   user.Controller
	}
)
