package generator

import (
	"strings"

	u "github.com/ravielze/fuzzy-broccoli/module/generator/utils"
	w "github.com/ravielze/fuzzy-broccoli/module/generator/wrapper"
)

func Regenerate(arg1, arg2 string) {
	packageName := strings.ToLower(arg1)
	cont, uc, repo := w.GetMethodWrapper(packageName)

	contContent := u.ReadFile(packageName, "controller.go")
	contContent += cont.String("(cont Controller)")
	u.WriteFile(packageName, "controller.go", contContent)

	ucContent := u.ReadFile(packageName, "usecase.go")
	ucContent += uc.String("(uc Usecase)")
	u.WriteFile(packageName, "usecase.go", ucContent)

	repoContent := u.ReadFile(packageName, "repository.go")
	repoContent += repo.String("(repo Repository)")
	u.WriteFile(packageName, "repository.go", repoContent)
}
