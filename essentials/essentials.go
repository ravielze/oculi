package essentials

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common/utils"
	"github.com/ravielze/oculi/middleware"
	"gorm.io/gorm"
)

type EssentialsModule struct {
}

func (EssentialsModule) Reset(db *gorm.DB) {}

func (EssentialsModule) Name() string {
	return "Essentials Module"
}

func NewModule(db *gorm.DB, g *gin.Engine) EssentialsModule {
	g.GET("/ping", func(ctx *gin.Context) {
		utils.OKAndResponseData(ctx, "Pong!")
	})
	g.GET("/reset", middleware.GetResetTokenMiddleware(), middleware.GetResetHandler(db))
	return EssentialsModule{}
}
