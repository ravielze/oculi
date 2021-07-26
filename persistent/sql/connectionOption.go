package sql

import (
	"time"

	"github.com/ravielze/oculi/logs"
)

type (
	ConnectionOption interface {
		Apply(o *ConnectionOptions)
	}
	ConnectionOptions struct {
		MaxIdleConnection, MaxOpenConnection int
		ConnMaxLifetime                      time.Duration
		LogMode                              bool
		logger                               logs.Logger
	}
)

func (o ConnectionOptions) Logger() logs.Logger {
	return o.logger
}

type withMaxIdleConnection int

func (w withMaxIdleConnection) Apply(o *ConnectionOptions) {
	o.MaxIdleConnection = int(w)
}

func WithMaxIdleConnection(maxIdleConnection int) ConnectionOption {
	return withMaxIdleConnection(maxIdleConnection)
}

type withMaxOpenConnection int

func (w withMaxOpenConnection) Apply(o *ConnectionOptions) {
	o.MaxOpenConnection = int(w)
}

func WithMaxOpenConnection(maxOpenConnection int) ConnectionOption {
	return withMaxOpenConnection(maxOpenConnection)
}

type withLogMode bool

func (w withLogMode) Apply(o *ConnectionOptions) {
	o.LogMode = bool(w)
}

func WithLogMode(logMode bool) ConnectionOption {
	return withLogMode(logMode)
}

type withConnMaxLifetime time.Duration

func (w withConnMaxLifetime) Apply(o *ConnectionOptions) {
	o.ConnMaxLifetime = time.Duration(w)
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) ConnectionOption {
	return withConnMaxLifetime(connMaxLifetime)
}

type withLogger struct{ logs.Logger }

func (w withLogger) Apply(o *ConnectionOptions) {
	o.logger = w.Logger
}

func WithLogger(logger logs.Logger) ConnectionOption {
	return withLogger{logger}
}
