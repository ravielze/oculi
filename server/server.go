package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/validator"
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
		Identifier() string
		ServerGracefullyDuration() time.Duration
		Logger() logs.Logger
		Validator() validator.Validator
		Close() error
	}

	Infrastructure interface {
		Register(ec *echo.Echo) error
		Health() echo.HandlerFunc
	}

	HookFunction func(res Resource) error
)
