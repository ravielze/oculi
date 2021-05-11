package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc IUsecase
}

func NewController(g *gin.Engine, uc IUsecase, a int64) IController {
	cont := Controller{
		uc: uc,
	}
	authGroup := g.Group("/auth")
	{
		authGroup.GET("/", func(ctx *gin.Context) {
			fmt.Println("Module auth.")
		})
	}
	return cont
}
