package external

import (
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/logs/zap"
	z "go.uber.org/zap"
)

func NewLogger(config *config.Env) (logs.Logger, error) {
	return zap.New(zap.Option{
		Level:  logs.GetLoggerLevel(config.LogLevel),
		Prefix: "",
	}, z.AddStacktrace(z.ErrorLevel), z.AddCallerSkip(1))
}
