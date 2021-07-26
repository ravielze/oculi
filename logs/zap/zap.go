package zap

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/ravielze/oculi/logs"
	"go.uber.org/zap"
)

type (
	Option struct {
		Level  log.Lvl
		Prefix string
	}
	logger struct {
		instance *zap.SugaredLogger
		prefix   string
		level    log.Lvl
	}
)

func (l logger) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (l logger) Output() io.Writer {
	return l
}

func (l *logger) SetOutput(w io.Writer) {

}

func (l logger) Prefix() string {
	return l.prefix
}

func (l *logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l logger) Level() log.Lvl {
	return l.level
}

func (l *logger) SetLevel(v log.Lvl) {
	l.level = v
}

func (l *logger) SetHeader(header string) {

}

func (l *logger) Info(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(args...)
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Infof(format, args...)
	}
}

func (l *logger) Infoj(j log.JSON) {
	if l.level != log.OFF {
		l.Infof("%+v\n", j)
	}
}

func (l *logger) Debug(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Debug(args...)
	}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Debugf(format, args...)
	}
}

func (l *logger) Debugj(j log.JSON) {
	if l.level != log.OFF {
		l.Debugf("%+v\n", j)
	}
}

func (l *logger) Error(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Error(args...)
	}
}

func (l *logger) Errorf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Errorf(format, args...)
	}
}

func (l *logger) Errorj(j log.JSON) {
	if l.level != log.OFF {
		l.Errorf("%+v\n", j)
	}
}

func (l *logger) Warning(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warn(args...)
	}
}

func (l *logger) Warningf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warnf(format, args...)
	}
}

func (l *logger) Warningj(j log.JSON) {
	if l.level != log.OFF {
		l.Warningf("%+v\n", j)
	}
}

func (l *logger) Fatal(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Fatal(args...)
	}
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Fatalf(format, args...)
	}
}

func (l *logger) Fatalj(j log.JSON) {
	if l.level != log.OFF {
		l.Fatalf("%+v\n", j)
	}
}

func (l *logger) Print(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(args...)
	}
}

func (l *logger) Println(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(args...)
	}
}

func (l *logger) Printf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Infof(format, args...)
	}
}

func (l *logger) Printj(j log.JSON) {
	if l.level != log.OFF {
		l.Printf("%+v\n", j)
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

func (l *logger) Panicj(j log.JSON) {
	if l.level != log.OFF {
		l.Panicf("%+v\n", j)
	}
}

func (l logger) Instance() interface{} {
	return l.instance
}

func (l logger) Log(msg string) {
	if l.level != log.OFF {
		l.instance.Info(msg)
	}
}

func DefaultLog() logs.Logger {
	logger, _ := New(true, Option{Level: log.INFO})
	return logger
}

func New(isDevelopment bool, logOption Option, options ...zap.Option) (logs.Logger, error) {
	var (
		instance *zap.Logger
		err      error
	)

	if isDevelopment {
		instance, err = zap.NewDevelopment(options...)
	} else {
		instance, err = zap.NewProduction(options...)
	}
	if err != nil {
		return nil, err
	}

	return &logger{
		instance: instance.Sugar(),
		level:    logOption.Level,
		prefix:   logOption.Prefix,
	}, nil
}

func (l *logger) logWithField(info logs.Info) *zap.SugaredLogger {
	logging := l.instance
	for key, message := range info.Data() {
		logging = logging.With(key, message)
	}
	return logging
}

func (l *logger) StandardPrint(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Info(info.Message())
	}
}

func (l *logger) StandardDebug(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Debug(info.Message())
	}
}

func (l *logger) StandardInfo(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Info(info.Message())
	}
}

func (l *logger) StandardWarn(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Warn(info.Message())
	}
}

func (l *logger) StandardError(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Error(info.Message())
	}
}

func (l *logger) StandardFatal(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Fatal(info.Message())
	}
}

func (l *logger) StandardPanic(info logs.Info) {
	if l.level != log.OFF {
		l.logWithField(info).Panic(info.Message())
	}
}
