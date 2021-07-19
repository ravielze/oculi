package infrastructures

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/example/infrastructures/rest"
	"go.uber.org/dig"
)

type (
	Component struct {
		dig.In

		Rest rest.Rest
	}
)

func (c Component) Register(ec *echo.Echo) error {
	return nil
}

func (c Component) Health() echo.HandlerFunc {
	return nil
}
