package domain

import (
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).End()
}
