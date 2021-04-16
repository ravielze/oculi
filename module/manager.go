package module

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/middleware"
	"github.com/ravielze/fuzzy-broccoli/module/auth"
	dvlp "github.com/ravielze/fuzzy-broccoli/module/development"
	"github.com/ravielze/fuzzy-broccoli/module/essentials"
	"github.com/ravielze/fuzzy-broccoli/module/filestorage"
	"gorm.io/gorm"
)

type Module interface {
	Name() string
}

var moduleList map[uint32]Module

func NewModule(db *gorm.DB, g *gin.Engine) map[uint32]Module {
	moduleList = map[uint32]Module{}
	moduleList[0] = dvlp.NewDevModule(db, g)
	moduleList[1] = essentials.NewEssentialsModule(db, g)
	moduleList[2] = auth.NewAuthModule(db, g)
	middleware.AuthModule = moduleList[2].(auth.AuthModule)
	moduleList[3] = filestorage.NewFileModule(db, g)
	return moduleList
}

func GetModule(moduleId uint32) Module {
	return moduleList[moduleId]
}
