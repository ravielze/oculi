package config

import "github.com/ravielze/oculi/config"

type (
	Env struct {
		ServiceName string `envconfig:"SERVICE_NAME" required:"true"`
		ServerPort  int    `envconfig:"SERVER_PORT" default:"8000" required:"true"`
	}
)

func NewConfig() (*Env, error) {
	conf := Env{}
	if err := config.New(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
