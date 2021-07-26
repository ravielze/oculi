package external

import (
	"github.com/labstack/echo/v4"
)

func NewEcho() (*echo.Echo, error) {
	return echo.New(), nil
}
