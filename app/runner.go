package app

import (
	"go.uber.org/dig"
)

type (
	ContainerConstructor func() (*dig.Container, error)
	Invoker              func(c *dig.Container) error
)

func Run(cc ContainerConstructor, inv Invoker) error {
	container, err := cc()
	if err != nil {
		return err
	}

	if err := inv(container); err != nil {
		return err
	}
	return nil
}
