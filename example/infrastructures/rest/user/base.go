package user

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
	g := ec.Group("/user")
	public := token.PublicEndpoint()
	g.POST("/login", c.Login, public)
	g.POST("/register", c.RegisterUser, public)
	return nil
}
