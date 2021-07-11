package rest

import (
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

		//Controllers
	}
)
