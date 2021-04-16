package module

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/module/auth"
	dvlp "github.com/ravielze/fuzzy-broccoli/module/development"
	"github.com/ravielze/fuzzy-broccoli/module/essentials"
	"gorm.io/gorm"
)

type Module interface {
	Name() string
}

func NewModule(db *gorm.DB, g *gin.Engine) map[uint32]Module {
	moduleList := map[uint32]Module{}
	moduleList[0] = essentials.NewEssentialsModule(db, g)
	moduleList[1] = dvlp.NewDevModule(db, g)
	moduleList[2] = auth.NewAuthModule(db, g)
	return moduleList
}
