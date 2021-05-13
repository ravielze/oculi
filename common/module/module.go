package module_manager

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

type Module interface {
	Name() string
	Reset(db *gorm.DB)
}

type ModuleList map[string]*Module

var moduleList ModuleList = ModuleList{}

func GetModule(moduleName string) Module {
	return *moduleList[moduleName]
}

func ResetAll(db *gorm.DB) {
	if db == nil {
		return
	}
	for _, module := range moduleList {
		(*module).Reset(db)
	}
	os.Exit(0)
}

func AddModule(m Module) {
	moduleList[m.Name()] = &m
}

func ShowModule() {
	if len(moduleList) > 0 {
		fmt.Println("| \u001b[44;1mOculi\u001b[0m | \033[33mModule:\033[0m")
		i := 1
		for _, x := range moduleList {
			fmt.Printf("| \u001b[44;1mOculi\u001b[0m | \033[34m[%d] \033[0m: \033[32m%s\033[0m\n", i, strings.Title((*x).Name()))
			i++
		}
	}
}
