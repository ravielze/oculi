package filemanager

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc IUsecase
}

func NewController(g *gin.Engine, uc IUsecase) IController {
	cont := Controller{
		uc: uc,
	}
	filemanagerGroup := g.Group("/filemanager")
	{
		filemanagerGroup.GET("/", func(ctx *gin.Context) {
			fmt.Println("Module filemanager.")
		})
	}
	return cont
}

func (cont Controller) GetFile(ctx *gin.Context) {
	panic("not implemented")
}

func (cont Controller) GetFilesByGroup(ctx *gin.Context) {
	panic("not implemented")
}
