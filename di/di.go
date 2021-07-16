package di

import (
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type (
	Registerable      func(c *dig.Container) error
	ContainerFunction func(once *sync.Once, container *dig.Container) (*dig.Container, error)
)

func Container(items []Registerable) ContainerFunction {
	return func(once *sync.Once, container *dig.Container) (*dig.Container, error) {
		var outer error
		once.Do(func() {
			container = dig.New()

			for i := range items {
				if err := items[i](container); err != nil {
					outer = err
					return
				}
			}
		})

		if outer != nil {
			return nil, errors.Wrap(outer, "cannot initialize container")
		}
		return container, nil
	}
}
