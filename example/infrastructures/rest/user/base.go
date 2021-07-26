package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/example/handlers"
	"github.com/ravielze/oculi/example/resources"
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
	g.POST("/login", c.Login)
	g.POST("/register", c.RegisterUser)
	g.GET("/check", c.Check)
	return nil
}
