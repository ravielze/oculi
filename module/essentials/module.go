package essentials

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	"github.com/ravielze/fuzzy-broccoli/middleware"
	"gorm.io/gorm"
)

type EssentialsModule struct {
}

func (EssentialsModule) Reset(db *gorm.DB) {}

func (EssentialsModule) Name() string {
	return "Essentials Module"
}

var ResetFunction func(db *gorm.DB)

func NewEssentialsModule(db *gorm.DB, g *gin.Engine) EssentialsModule {
	g.GET("/ping", func(ctx *gin.Context) {
		utils.OKAndResponseData(ctx, "Pong!")
	})
	g.GET("/reset", middleware.GetResetTokenMiddleware(), func(ctx *gin.Context) {
		ResetFunction(db)
	})
	return EssentialsModule{}
}
