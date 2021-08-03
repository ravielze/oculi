package di

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type (
	registrant struct {
		container *dig.Container
		err       error
	}

	Registrant interface {
		Register(rf Registerable) Registrant
		Provide(constructor interface{}, opts ...dig.ProvideOption) Registrant
		Proceed() error
	}
)

func NewRegistrant(container *dig.Container) Registrant {
	return &registrant{
		container: container,
		err:       nil,
	}
}

func SimpleRegistrant(container *dig.Container, constructor interface{}, opts ...dig.ProvideOption) error {
	return NewRegistrant(container).Provide(constructor, opts...).Proceed()
}

func (r *registrant) Provide(constructor interface{}, opts ...dig.ProvideOption) Registrant {
	if r.err == nil {
		if err := r.container.Provide(constructor, opts...); err != nil {
			r.err = errors.Wrap(err, "error on initialize")
		}
	}
	return r
}

func (r *registrant) Register(rf Registerable) Registrant {
	if r.err == nil {
		if err := rf(r.container); err != nil {
			r.err = err
		}
	}
	return r
}

func (r *registrant) Proceed() error {
	return r.err
}
