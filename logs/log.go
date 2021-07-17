package logs

import (
	"io"

	"github.com/labstack/gommon/log"
)

type (
	Logger interface {
		Output() io.Writer
		SetOutput(w io.Writer)
		Prefix() string
		SetPrefix(p string)
		Level() log.Lvl
		SetLevel(v log.Lvl)
		SetHeader(h string)
		Print(i ...interface{})
		Printf(format string, args ...interface{})
		Printj(j log.JSON)
		Debug(i ...interface{})
		Debugf(format string, args ...interface{})
		Debugj(j log.JSON)
		Info(i ...interface{})
		Infof(format string, args ...interface{})
		Infoj(j log.JSON)
		Warn(i ...interface{})
		Warnf(format string, args ...interface{})
		Warnj(j log.JSON)
		Error(i ...interface{})
		Errorf(format string, args ...interface{})
		Errorj(j log.JSON)
		Fatal(i ...interface{})
		Fatalj(j log.JSON)
		Fatalf(format string, args ...interface{})
		Panic(i ...interface{})
		Panicj(j log.JSON)
		Panicf(format string, args ...interface{})
	}
)

func GetLoggerLevel(level string) log.Lvl {
	switch level {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	default:
		return log.INFO
	}
}
