package resources

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/docs"
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/response"
	"github.com/ravielze/oculi/validator"
	"go.uber.org/dig"
)

var (
	once       sync.Once
	identifier string
	uptime     time.Time
)

type (
	Resource struct {
		dig.In

		EchoData      *echo.Echo
		Log           logs.Logger
		Responder     response.Responder
		Config        *config.Env
		ValidatorData validator.Validator
		Database      sql.API
		Documentation docs.Documentation
	}
)

func (r Resource) Echo() *echo.Echo {
	return r.EchoData
}
func (r Resource) ServiceName() string {
	return r.Config.ServiceName
}
func (r Resource) ServerPort() int {
	return r.Config.ServerPort
}

func (r Resource) Identifier() string {
	once.Do(func() {
		uptime = time.Now()
		if identifier == "" {
			identifier = r.Config.ServiceName + " " + uptime.String()
		}
	})
	return identifier
}

func (r Resource) Uptime() time.Time {
	once.Do(func() {
		uptime = time.Now()
		if identifier == "" {
			identifier = r.Config.ServiceName + " " + uptime.String()
		}
	})
	return uptime
}

func (r Resource) ServerGracefullyDuration() time.Duration {
	return r.Config.GracefullyDuration
}

func (r Resource) Logger() logs.Logger {
	return r.Log
}

func (r Resource) Validator() validator.Validator {
	return r.ValidatorData
}

func (r Resource) Close() error {
	var errMessage = make([]string, 0)

	if err := r.EchoData.Close(); err != nil {
		errMessage = append(errMessage, err.Error())
	}

	if len(errMessage) > 0 {
		return errors.New(strings.Join(errMessage, "\n"))
	}
	return nil
}
