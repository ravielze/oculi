package generator

import (
	_ "embed"
	"fmt"

	u "github.com/ravielze/oculi/generator/utils"
)

//go:embed template/main.txt
var mainContent string

func Init() {

	if u.IsPackageExist("app") {
		fmt.Println("That package is already exist.")
		return
	}
	fmt.Printf("Initiating oculi project...\n")

	u.WriteFile("app", "main.go", mainContent)
}
