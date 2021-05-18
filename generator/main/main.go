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
	case "add", "a":
		CheckArgs(2)
		generator.Generate(os.Args[2], os.Args[3])
	case "add-simple", "add-s", "as":
		CheckArgs(1)
		generator.GenerateSimple(os.Args[2])
	case "update", "upd", "u":
		CheckArgs(2)
		generator.Regenerate(os.Args[2], os.Args[3])
	case "init", "i":
		CheckArgs(0)
		generator.Init()
	case "preset":
		generator.GenerateAuthPreset()
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
			case "add", "a":
				fmt.Printf("Usage: %s [add|a] [packageName] [moduleName]\n", os.Args[0])
			case "add-simple", "add-s", "as":
				fmt.Printf("Usage: %s [add-simple|add-s|as] [packageName]\n", os.Args[0])
			case "update", "upd", "u":
				fmt.Printf("Usage: %s [update|upd|u] [packageName] [moduleName]\n", os.Args[0])
			}
		}
		panic("Invalid command.")
	}
}

func ShowHelp() {
	cmd := os.Args[0]
	fmt.Println(cmd, "help")
	fmt.Println(cmd, "[init|i]")
	fmt.Println(cmd, "[add|a] [packageName] [moduleName]")
	fmt.Println(cmd, "[add-simple|add-s|as] [packageName]")
	fmt.Println(cmd, "[update|upd|u] [packageName] [moduleName]")
}
