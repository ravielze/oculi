package di

import (
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type (
	Registerable      func(c *dig.Container) error
	ContainerFunction func(once *sync.Once, container *dig.Container) (*dig.Container, error)

	Registrant struct {
		container *dig.Container
		err       error
	}
)

func NewRegistrant(container *dig.Container) *Registrant {
	return &Registrant{
		container: container,
		err:       nil,
	}
}

func (r *Registrant) Provide(constructor interface{}, opts ...dig.ProvideOption) *Registrant {
	if r.err == nil {
		if err := r.container.Provide(constructor, opts...); err != nil {
			r.err = errors.Wrap(err, "error on initialize")
		}
	}
	return r
}

func (r *Registrant) Register(rf Registerable) *Registrant {
	if r.err == nil {
		if err := rf(r.container); err != nil {
			r.err = err
		}
	}
	return r
}

func (r *Registrant) End() error {
	return r.err
}

func GenerateContainerFunction(items []Registerable) ContainerFunction {
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
