package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/example/config"
	"go.uber.org/dig"
)

type (
	Resource struct {
		dig.In

		GinEngine *gin.Engine
		Config    *config.Env
	}
)

func (r Resource) Gin() *gin.Engine {
	return r.GinEngine
}
func (r Resource) ServiceName() string {
	return r.Config.ServiceName
}
func (r Resource) ServerPort() int {
	return r.Config.ServerPort
}
func (r Resource) Close() error {
	return nil
}
