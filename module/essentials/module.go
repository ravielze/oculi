package essentials

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	"gorm.io/gorm"
)

type EssentialsModule struct {
}

func (EssentialsModule) Name() string {
	return "Essentials Module"
}

func NewEssentialsModule(db *gorm.DB, g *gin.Engine) EssentialsModule {
	g.GET("/ping", func(ctx *gin.Context) {
		utils.OKAndResponseData(ctx, "Pong!")
	})
	return EssentialsModule{}
}
