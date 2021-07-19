package config

import (
	"time"

	"github.com/ravielze/oculi/config"
)

type (
	Env struct {
		ServiceName        string        `envconfig:"SERVICE_NAME" required:"true"`
		ServerPort         int           `envconfig:"SERVER_PORT" default:"8000" required:"true"`
		GracefullyDuration time.Duration `envconfig:"GRACEFULLY_DURATION" default:"5s"`

		LogLevel string `envconfig:"LOG_LEVEL" default:"INFO" required:"true"`
	}
)

func NewConfig() (*Env, error) {
	conf := Env{}
	if err := config.New(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
