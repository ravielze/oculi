package generator

import (
	_ "embed"
	"fmt"

	u "github.com/ravielze/oculi/generator/utils"
)

//go:embed template/main.txt
var mainContent string

func Init() {

	fmt.Printf("Initiating oculi project...\n")

	u.WriteFile("app", "main.go", mainContent)
}
