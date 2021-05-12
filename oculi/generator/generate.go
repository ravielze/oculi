package generator

import (
	_ "embed"
	"fmt"
	"strings"

	u "github.com/ravielze/oculi/oculi/generator/utils"
)

//go:embed template/entity.txt
var entityRawContent string

//go:embed template/repo.txt
var repositoryRawContent string

//go:embed template/uc.txt
var usecaseRawContent string

//go:embed template/cont.txt
var controllerRawContent string

//go:embed template/module.txt
var moduleRawContent string

func MakePlaceholders(packageName, moduleName string) []u.Placeholder {
	return []u.Placeholder{
		u.NewPlaceholder("$$package$$", packageName),
		u.NewPlaceholder("$$module$$", moduleName),
		u.NewPlaceholder("$$module_lower$$", strings.ToLower(moduleName)),
	}
}

func Generate(arg1, arg2 string) {
	packageName := strings.ToLower(arg1)
	moduleName := strings.Title(strings.ToLower(arg2))
	moduleNameLower := strings.ToLower(arg2)

	placeholders := MakePlaceholders(packageName, moduleName)

	fmt.Printf("Generating package %s: entity -> %s\n", packageName, moduleName)

	entityContent := u.Replacer(entityRawContent, placeholders)
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 1, moduleNameLower, "entity.go"), entityContent)

	repositoryContent := u.Replacer(repositoryRawContent, placeholders)
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 2, moduleNameLower, "repository.go"), repositoryContent)

	usecaseContent := u.Replacer(usecaseRawContent, placeholders)
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 3, moduleNameLower, "usecase.go"), usecaseContent)

	controllerContent := u.Replacer(controllerRawContent, placeholders)
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 4, moduleNameLower, "controller.go"), controllerContent)

	moduleContent := u.Replacer(moduleRawContent, placeholders)
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 0, moduleNameLower, "module.go"), moduleContent)
}
