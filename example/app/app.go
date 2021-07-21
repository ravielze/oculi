package app

import (
	"github.com/ravielze/oculi/app"
	"github.com/ravielze/oculi/example/di"
	"github.com/ravielze/oculi/example/infrastructures"
	"github.com/ravielze/oculi/example/resources"

	webserver "github.com/ravielze/oculi/server/echo"
	"go.uber.org/dig"
)

func Run() {
	invoker := func(container *dig.Container) error {
		return container.Invoke(func(i infrastructures.Component, r resources.Resource) error {
			s := webserver.New(i, r)
			if r.Config.IsDevelopment() {
				s.DevelopmentMode()
			}
			return s.Run()
		})
	}

	if err := app.Run(di.Container, invoker); err != nil {
		panic(err)
	}
}
