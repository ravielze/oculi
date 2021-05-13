package main

import (
	"fmt"
	"os"

	"github.com/ravielze/oculi/generator"
)

func main() {

	defer handleError()
	if len(os.Args) <= 1 {
		ShowHelp()
		return
	}
	switch os.Args[1] {
	case "help", "h":
		ShowHelp()
	case "addmodule", "am":
		CheckArgs(2)
		generator.Generate(os.Args[2], os.Args[3])
	case "updatemodule", "um":
		CheckArgs(2)
		generator.Regenerate(os.Args[2], os.Args[3])
	case "init", "i":
		CheckArgs(0)
		generator.Init()
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
			case "init", "i":
				fmt.Printf("Usage: %s [init|i]\n", os.Args[0])
			case "addmodule", "am":
				fmt.Printf("Usage: %s [addmodule|am] [packageName] [moduleName]\n", os.Args[0])
			case "updatemodule", "um":
				fmt.Printf("Usage: %s [updatemodule|um] [packageName] [moduleName]\n", os.Args[0])
			}
		}
		panic("Invalid command.")
	}
}

func ShowHelp() {
	cmd := os.Args[0]
	fmt.Println(cmd, "help")
	fmt.Println(cmd, "[init|i]")
	fmt.Println(cmd, "[addmodule|am] [packageName] [moduleName]")
	fmt.Println(cmd, "[updatemodule|um] [packageName] [moduleName]")
}
