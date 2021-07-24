package infrastructures

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ravielze/oculi/example/infrastructures/rest"
	oculiRest "github.com/ravielze/oculi/infrastructures/rest"
	ws "github.com/ravielze/oculi/server/echo"
	"go.uber.org/dig"
)

type (
	Component struct {
		dig.In

		Rest rest.Rest
	}
)

func (c Component) Register(ec *echo.Echo) error {
	ec.Pre(middleware.RemoveTrailingSlash())
	v1 := ec.Group("/v1")
	if c.Rest.Resource.Config.IsDevelopment() {
		defer ws.InfoRoutes(ec)()
	}
	return oculiRest.Register(v1, &c.Rest.Controller)
}

func (c Component) Health() echo.HandlerFunc {
	return c.Rest.Controller.Health.Check
}
