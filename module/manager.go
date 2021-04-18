package module

import (
	"os"

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
	Reset(*gorm.DB)
}

var moduleList map[uint32]Module

func NewModule(db *gorm.DB, g *gin.Engine) map[uint32]Module {
	middleware.InstallCors(g)
	middleware.InstallGZipCompressor(g)
	middleware.InstallDefaultLimiter(g)
	moduleList = map[uint32]Module{}
	moduleList[0] = dvlp.NewDevModule(db, g)

	essentialsModule := essentials.NewEssentialsModule(db, g)
	essentials.ResetFunction = ResetAll
	moduleList[1] = essentialsModule

	authModule := auth.NewAuthModule(db, g)
	middleware.AuthModule = &authModule
	moduleList[2] = authModule

	moduleList[3] = filestorage.NewFileModule(db, g)
	return moduleList
}

func GetModule(moduleId uint32) Module {
	return moduleList[moduleId]
}

func ResetAll(db *gorm.DB) {
	if db == nil {
		return
	}
	for _, module := range moduleList {
		module.Reset(db)
	}
	os.Exit(0)
}
