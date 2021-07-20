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

		DatabaseAddress           string        `envconfig:"DB_ADDRESS" required:"true"`
		DatabaseUsername          string        `envconfig:"DB_USERNAME" required:"true"`
		DatabasePassword          string        `envconfig:"DB_PASSWORD" required:"true"`
		DatabaseName              string        `envconfig:"DB_NAME" required:"true"`
		DatabaseMaxIdleConnection int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
		DatabaseMaxOpenConnection int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"25"`
		DatabaseConnMaxLifetime   time.Duration `envconfig:"DB_CONNECTION_MAX_LIFE_TIME" default:"60s"`
		DatabaseLogMode           bool          `envconfig:"DB_LOG_MODE" default:"false"`
	}
)

func NewConfig() (*Env, error) {
	conf := Env{}
	if err := config.New(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
