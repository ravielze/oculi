package generator

import (
	_ "embed"
	"fmt"
	"strings"

	u "github.com/ravielze/oculi/generator/utils"
)

//go:embed template/module/entity.txt
var entityRawContent string

//go:embed template/module/repo.txt
var repositoryRawContent string

//go:embed template/module/uc.txt
var usecaseRawContent string

//go:embed template/module/cont.txt
var controllerRawContent string

//go:embed template/module/module.txt
var moduleRawContent string

//go:embed template/module/simplemodule.txt
var simpleModuleRawContent string

func MakePlaceholders(packageName, moduleName string) []u.Placeholder {
	return []u.Placeholder{
		u.NewPlaceholder("$$package$$", packageName),
		u.NewPlaceholder("$$module$$", moduleName),
		u.NewPlaceholder("$$module_lower$$", strings.ToLower(moduleName)),
	}
}

func Generate(arg1, arg2 string) {
	packageName := strings.ToLower(arg1)

	if u.IsPackageExist(packageName) {
		fmt.Println("That package is already exist.")
		return
	}

	moduleName := strings.Title(strings.ToLower(arg2))
	moduleNameLower := strings.ToLower(arg2)

	placeholders := MakePlaceholders(packageName, moduleName)

	fmt.Printf("Generating package %s: entity -> %s\n", packageName, moduleName)

	u.ReplacerWriter(entityRawContent,
		packageName,
		fmt.Sprintf("%d_%s_%s", 1, moduleNameLower, "entity.go"),
		placeholders,
	)
	u.ReplacerWriter(repositoryRawContent,
		packageName,
		fmt.Sprintf("%d_%s_%s", 2, moduleNameLower, "repository.go"),
		placeholders,
	)
	u.ReplacerWriter(usecaseRawContent,
		packageName,
		fmt.Sprintf("%d_%s_%s", 3, moduleNameLower, "usecase.go"),
		placeholders,
	)
	u.ReplacerWriter(controllerRawContent,
		packageName,
		fmt.Sprintf("%d_%s_%s", 4, moduleNameLower, "controller.go"),
		placeholders,
	)
	u.ReplacerWriter(moduleRawContent,
		packageName,
		fmt.Sprintf("%d_%s_%s", 0, moduleNameLower, "module.go"),
		placeholders,
	)
}

func GenerateSimple(arg1 string) {
	packageName := strings.ToLower(arg1)

	if u.IsPackageExist(packageName) {
		fmt.Println("That package is already exist.")
		return
	}

	placeholders := MakePlaceholders(packageName, "")

	fmt.Printf("Generating package %s: simple mode\n", packageName)

	simpleModuleContent := u.Replacer(simpleModuleRawContent, placeholders)
	u.WriteFile(packageName, packageName+".go", simpleModuleContent)
}
