package main

import (
	"fmt"
	"os"

	"github.com/ravielze/oculi/generator"
)

func main() {

	defer handleError()
	CheckArgs(0)
	switch os.Args[1] {
	case "help", "h":
		ShowHelp()
	case "addmodule", "am":
		CheckArgs(2)
		generator.Generate(os.Args[2], os.Args[3])
	case "updatemodule", "um":
		CheckArgs(2)
		generator.Regenerate(os.Args[2], os.Args[3])
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
