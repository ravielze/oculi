package generator

import (
	"fmt"
	"strings"

	u "github.com/ravielze/oculi/generator/utils"
	w "github.com/ravielze/oculi/generator/wrapper"
)

func Regenerate(arg1, arg2 string) {
	packageName := strings.ToLower(arg1)
	moduleNameLower := strings.ToLower(arg2)

	if !u.IsPackageExist(packageName) {
		fmt.Println("That package is not exist.")
		return
	}
	cont, uc, repo := w.GetMethodWrapper(packageName, moduleNameLower)

	contContent := u.ReadFile(packageName, fmt.Sprintf("%d_%s_%s", 4, moduleNameLower, "controller.go"))
	contContent += cont.String("(cont Controller)")
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 4, moduleNameLower, "controller.go"), contContent)

	ucContent := u.ReadFile(packageName, fmt.Sprintf("%d_%s_%s", 3, moduleNameLower, "usecase.go"))
	ucContent += uc.String("(uc Usecase)")
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 3, moduleNameLower, "usecase.go"), ucContent)

	repoContent := u.ReadFile(packageName, fmt.Sprintf("%d_%s_%s", 2, moduleNameLower, "repository.go"))
	repoContent += repo.String("(repo Repository)")
	u.WriteFile(packageName, fmt.Sprintf("%d_%s_%s", 2, moduleNameLower, "repository.go"), repoContent)
}
