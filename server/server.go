package server

import (
	"time"

	"github.com/gin-gonic/gin"
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
		Gin() *gin.Engine
		ServiceName() string
		ServerPort() int
		ServerGracefullyDuration() time.Duration
		Close() error
	}

	Infrastructure interface {
		Register(gin *gin.Engine) error
		Health() gin.HandlerFunc
	}

	HookFunction func(res Resource) error
)
