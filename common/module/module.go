package module

import (
	"os"

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

func AddModule(m Module) {
	moduleList[m.Name()] = m
}
