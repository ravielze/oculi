package essentials

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common/middleware"
	"github.com/ravielze/oculi/common/utils"
	"gorm.io/gorm"
)

type Module struct {
}

func (Module) Reset(db *gorm.DB) {}

func (Module) Name() string {
	return "essentials"
}

func NewModule(db *gorm.DB, g *gin.Engine) Module {
	g.GET("/ping", func(ctx *gin.Context) {
		utils.OKAndResponseData(ctx, "Pong!")
	})
	g.GET("/reset", middleware.GetResetTokenMiddleware(), middleware.GetResetHandler(db))
	return Module{}
}
