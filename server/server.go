package server

import (
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Server interface {
		Run() error

		BeforeRun(hf HookFunction) Server
		AfterRun(hf HookFunction) Server
		BeforeExit(hf HookFunction) Server
		AfterExit(hf HookFunction) Server
	}

	Resource interface {
		Echo() *echo.Echo
		ServiceName() string
		ServerPort() int
		ServerGracefullyDuration() time.Duration
		Logger() interface{}    //TODO
		Validator() interface{} //TODO
		Close() error
	}

	Infrastructure interface {
		Register(ec *echo.Echo) error
		Health() echo.HandlerFunc
		HealthRoutes() echo.HandlerFunc
	}

	HookFunction func(res Resource) error
)
