package health

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/example/handlers"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/middleware/token"
	"go.uber.org/dig"
)

type (
	Controller struct {
		dig.In

		Handler  handlers.Handler
		Resource resources.Resource
	}
)

func (c Controller) Register(ec *echo.Group) error {
	g := c.Resource.Echo()
	public := token.PublicEndpoint()
	g.DELETE("/reset", c.Reset, public)
	return nil
}
