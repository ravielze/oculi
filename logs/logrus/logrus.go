package logrus

import (
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/ravielze/oculi/logs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type (
	format string

	Option struct {
		Level       log.Lvl
		LogFilePath string
		Formatter   format
		Prefix      string
	}

	logger struct {
		instance *logrus.Logger
		level    log.Lvl
		prefix   string
	}
)

const (
	JSON format = "JSON"
	TEXT format = "TEXT"
)

func New(option *Option) (logs.Logger, error) {
	instance := logrus.New()

	switch option.Level {
	case log.INFO:
		instance.Level = logrus.InfoLevel
	case log.DEBUG:
		instance.Level = logrus.DebugLevel
	case log.WARN:
		instance.Level = logrus.WarnLevel
	case log.ERROR:
		instance.Level = logrus.ErrorLevel
	default:
		instance.Level = logrus.ErrorLevel
	}

	var formatter logrus.Formatter
	switch option.Formatter {
	case JSON:
		formatter = &logrus.JSONFormatter{}
	default:
		formatter = &logrus.TextFormatter{}
	}

	instance.Formatter = formatter
	if option.LogFilePath != "" {
		if _, err := os.Stat(option.LogFilePath); os.IsNotExist(err) {
			if _, err = os.Create(option.LogFilePath); err != nil {
				return nil, errors.Wrapf(err, "failed to create log file %s", option.LogFilePath)
			}
		}
		maps := lfshook.PathMap{
			logrus.InfoLevel:  option.LogFilePath,
			logrus.DebugLevel: option.LogFilePath,
			logrus.ErrorLevel: option.LogFilePath,
		}
		instance.Hooks.Add(lfshook.NewHook(maps, formatter))
	}

	return &logger{
		instance: instance,
		level:    option.Level,
		prefix:   option.Prefix,
	}, nil
}

func (l *logger) Output() io.Writer {
	return l.instance.Out
}

func (l *logger) Prefix() string {
	return l.prefix
}

func (l *logger) Level() log.Lvl {
	return l.level
}

func (l *logger) Print(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Print(i...)
	}
}

func (l *logger) Println(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Println(i...)
	}
}

func (l *logger) Printf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Printf(format, i...)
	}
}

func (l *logger) Printj(j log.JSON) {
	if l.level != log.OFF {
		l.Printf("%+v\n", j)
	}
}

func (l *logger) Debug(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Debug(i...)
	}
}

func (l *logger) Debugf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Debugf(format, i...)
	}
}

func (l *logger) Debugj(j log.JSON) {
	if l.level != log.OFF {
		l.Debugf("%+v\n", j)
	}
}

func (l *logger) Info(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(i...)
	}
}

func (l *logger) Infof(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Infof(format, i...)
	}
}

func (l *logger) Infoj(j log.JSON) {
	if l.level != log.OFF {
		l.Infof("%+v\n", j)
	}
}

func (l *logger) Warn(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warn(i...)
	}
}

func (l *logger) Warnf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warnf(format, i...)
	}
}

func (l *logger) Warnj(j log.JSON) {
	if l.level != log.OFF {
		l.Warnf("%+v\n", j)
	}
}

func (l *logger) Error(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Error(i...)
	}
}

func (l *logger) Errorf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Errorf(format, i...)
	}
}

func (l *logger) Errorj(j log.JSON) {
	if l.level != log.OFF {
		l.Errorf("%+v\n", j)
	}
}

func (l *logger) Fatal(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Fatal(i...)
	}
}

func (l *logger) Fatalf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Fatalf(format, i...)
	}
}

func (l *logger) Fatalj(j log.JSON) {
	if l.level != log.OFF {
		l.Fatalf("%+v\n", j)
	}
}

func (l *logger) Panic(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Panic(i...)
	}
}

func (l *logger) Panicf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Panicf(format, i...)
	}
}

func (l *logger) SetHeader(h string) {}

func (l *logger) Panicj(j log.JSON) {
	if l.level != log.OFF {
		l.Panicf("%+v\n", j)
	}
}

func (l *logger) Instance() interface{} {
	return l.instance
}

func (l *logger) Log(msg string) {
	if l.level != log.OFF {
		l.instance.Info(msg)
	}
}

func (l *logger) SetLevel(v log.Lvl) {
	l.level = v
	l.instance.SetLevel(getLevel(v))
}

func getLevel(lvl log.Lvl) logrus.Level {
	switch lvl {
	case log.INFO:
		return logrus.InfoLevel
	case log.DEBUG:
		return logrus.DebugLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	default:
		return logrus.ErrorLevel
	}
}

func (l *logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *logger) SetOutput(w io.Writer) {
	l.instance.Out = w
}
