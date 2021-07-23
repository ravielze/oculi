package resources

import (
	"github.com/ravielze/oculi/di"
	"github.com/ravielze/oculi/example/resources/external"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Provide(external.NewEcho).
		Provide(external.NewValidator).
		Provide(external.NewResponder).
		Provide(external.NewLogger).
		Provide(external.NewDocs).
		Provide(external.NewPostgreSQL).
		Provide(external.NewTokenizer).
		Provide(external.NewHash).
		End()
}
