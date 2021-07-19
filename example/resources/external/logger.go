package external

import (
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/logs/zap"
	z "go.uber.org/zap"
)

func NewLogger(config *config.Env) (logs.Logger, error) {
	option := z.AddStacktrace(z.ErrorLevel)

	return zap.New(zap.Option{
		Level:  logs.GetLoggerLevel(config.LogLevel),
		Prefix: "",
	}, option)
}
