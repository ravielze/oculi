package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func New(object interface{}) error {
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", object); err != nil {
			return errors.Wrap(err, "failed to read from env variable")
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", object); err != nil {
		return errors.Wrap(err, "failed to read from env variable")
	}

	return nil
}
