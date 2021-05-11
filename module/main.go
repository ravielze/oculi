package main

import (
	"fmt"
	"os"
	"strings"

	generator_utils "github.com/ravielze/fuzzy-broccoli/module/generator/utils"
	w "github.com/ravielze/fuzzy-broccoli/module/generator/wrapper"

	"github.com/ravielze/fuzzy-broccoli/module/generator"
)

func main() {

	defer handleError()
	CheckArgs(0)
	switch os.Args[1] {
	case "help", "h":
		ShowHelp()
	case "gen", "generate", "g":
		CheckArgs(2)
		generator.Generate(os.Args[2], os.Args[3])
	case "regenerate", "regen", "r":
		CheckArgs(2)
		Testing(os.Args[2], os.Args[3])
	default:
		fmt.Println("Command not found. Try", os.Args[0], "help")
	}
}

func handleError() {
	if a := recover(); a != nil {
		fmt.Println("Error:", a)
	}
}
func CheckArgs(argsNeeded int) {
	if len(os.Args) < (argsNeeded + 2) {
		if len(os.Args) >= 2 {
			switch os.Args[1] {
			case "gen", "generate", "g":
				fmt.Printf("Usage: %s [generate|gen|g] [packageName] [moduleName]\n", os.Args[0])
			case "regenerate", "regen", "r":
				fmt.Printf("Usage: %s [regenerate|regen|r] [packageName] [moduleName]\n", os.Args[0])
			}
		}
		panic("Invalid command.")
	}
}

func ShowHelp() {
	cmd := os.Args[0]
	fmt.Println(cmd, "help")
	fmt.Println(cmd, "[generate|gen|g] [packageName] [moduleName]")
	fmt.Println(cmd, "[regenerate|regen|r] [packageName] [moduleName]")
}

func GetFileContent(folderName, fileName, content string) {
	//TODO
}

func Testing(arg1, arg2 string) []string {
	packageName := strings.ToLower(arg1)
	x := w.GetMethodWrapper(packageName)["Repository"].String("(repo Repository)")
	repo := generator_utils.ReadFile(packageName, "repository.go")
	repo += x
	generator_utils.WriteFile(packageName, "repository.go", repo)
	// fmt.Println(generator.ReadDeclaredFunctions(packageName, "controller.go"))
	// fmt.Println(generator.ReadDeclaredFunctions(packageName, "repository.go"))
	// fmt.Println(generator.ReadDeclaredFunctions(packageName, "usecase.go"))
	// fmt.Println()
	//generator.ReadDeclaredInterfaces(packageName, "entity.go")
	return nil
}
