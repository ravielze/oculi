package di

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type (
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

func SimpleRegistrant(container *dig.Container, constructor interface{}, opts ...dig.ProvideOption) error {
	return NewRegistrant(container).Provide(constructor, opts...).End()
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
