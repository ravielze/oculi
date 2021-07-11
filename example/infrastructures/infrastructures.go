package infrastructures

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/example/infrastructures/rest"
	"go.uber.org/dig"
)

type (
	Component struct {
		dig.In

		Rest rest.Rest
	}
)

func (c Component) Register(gin *gin.Engine) error {
	return nil
}

func (c Component) Health() gin.HandlerFunc {
	return nil
}
