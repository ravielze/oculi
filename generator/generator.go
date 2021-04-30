package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed module/entity.txt
var entityRawContent string

//go:embed module/repo.txt
var repositoryRawContent string

//go:embed module/uc.txt
var usecaseRawContent string

//go:embed module/cont.txt
var controllerRawContent string

var packageName string = "module"
var moduleName string = "Entity"

func main() {

	defer handleError()
	checkArgs(os.Args)
	packageName = strings.ToLower(os.Args[1])
	moduleName = strings.Title(strings.ToLower(os.Args[2]))
	moduleNameLower := strings.ToLower(moduleName)
	fmt.Printf("Generating package %s: entity -> %s\n", packageName, moduleName)
	entityContent := strings.ReplaceAll(entityRawContent, "$$package$$", packageName)
	entityContent = strings.ReplaceAll(entityContent, "$$module$$", moduleName)
	entityContent = strings.ReplaceAll(entityContent, "$$module_lower$$", moduleNameLower)
	WriteFile(moduleNameLower, "entity.go", entityContent)

	repositoryContent := strings.ReplaceAll(repositoryRawContent, "$$package$$", packageName)
	WriteFile(moduleNameLower, "repository.go", repositoryContent)

	usecaseContent := strings.ReplaceAll(usecaseRawContent, "$$package$$", packageName)
	WriteFile(moduleNameLower, "usecase.go", usecaseContent)

	controllerContent := strings.ReplaceAll(controllerRawContent, "$$package$$", packageName)
	WriteFile(moduleNameLower, "controller.go", controllerContent)
}

func handleError() {
	if a := recover(); a != nil {
		fmt.Println("Error: ", a)
	}
}

func ShowHelp() {
	//TODO help
}

func GetFileContent(folderName, fileName, content string) {
	//TODO
}

func GetAllFunction(folderName, fileName string) []string {
	//TODO
	return nil
}

func checkArgs(args []string) {
	if len(args) < 3 {
		//TODO help
		panic(fmt.Sprintf("Usage: %s [packageName] [moduleName]", args[0]))
	}
}

func WriteFile(folderName, fileName, content string) {
	if _, err0 := os.Stat(folderName); os.IsNotExist(err0) {
		os.Mkdir(folderName, 0755)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", folderName, fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(content)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("Writing %s/%s\n", folderName, fileName)
}
