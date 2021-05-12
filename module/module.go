package module

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/essentials"
	"github.com/ravielze/oculi/middleware"
	"gorm.io/gorm"
)

type Module interface {
	Name() string
	Reset(*gorm.DB)
}

type ModuleList map[string]Module

var moduleList ModuleList = ModuleList{}

func GetModule(moduleName string) Module {
	return moduleList[moduleName]
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

func NewModule(db *gorm.DB, g *gin.Engine, devMode bool) *ModuleList {
	middleware.InstallCors(g)
	middleware.InstallDefaultLimiter(g)

	essentialsModule := essentials.NewModule(db, g)

	AddModule(essentialsModule)
	return &moduleList
}

func AddModule(m Module) {
	moduleList[m.Name()] = m
}
