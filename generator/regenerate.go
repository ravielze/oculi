package generator

import (
	"fmt"
	"strings"

	u "github.com/ravielze/oculi/generator/utils"
	w "github.com/ravielze/oculi/generator/wrapper"
)

func Regenerate(arg1, arg2 string) {
	packageName := strings.ToLower(arg1)

	if !u.IsPackageExist(packageName) {
		fmt.Println("That package is not exist.")
		return
	}
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
