package di

import (
	"sync"

	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/example/domain"
	"github.com/ravielze/oculi/example/handlers"
	"github.com/ravielze/oculi/example/resources"
	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	items := []di.Registerable{
		config.Register,
		resources.Register,
		domain.Register,
		handlers.Register,
	}
	return di.Container(items)(&once, container)
}
