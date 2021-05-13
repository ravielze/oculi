package module

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module interface {
	Name() string
	Reset(db *gorm.DB)
	NewModule(db *gorm.DB, g *gin.Engine, obj ...interface{})
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

func AddModule(m Module) {
	moduleList[m.Name()] = m
}
