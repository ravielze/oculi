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

	placeholders := MakePlaceholders(packageName, moduleName)

	fmt.Printf("Generating package %s: entity -> %s\n", packageName, moduleName)

	entityContent := u.Replacer(entityRawContent, placeholders)
	u.WriteFile(packageName, "entity.go", entityContent)

	repositoryContent := u.Replacer(repositoryRawContent, placeholders)
	u.WriteFile(packageName, "repository.go", repositoryContent)

	usecaseContent := u.Replacer(usecaseRawContent, placeholders)
	u.WriteFile(packageName, "usecase.go", usecaseContent)

	controllerContent := u.Replacer(controllerRawContent, placeholders)
	u.WriteFile(packageName, "controller.go", controllerContent)
}
