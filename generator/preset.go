package generator

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"

	u "github.com/ravielze/oculi/generator/utils"
)

//go:embed template/preset/auth/*
var authFiles embed.FS

//go:embed template/preset/filemanager/*
var fileManagerFiles embed.FS

func GeneratePresets(arg1 string) {
	switch strings.ToLower(arg1) {
	case "auth":
		generateAuthPreset()
	case "filemanager", "file":
		generateFileManagerPreset()
	case "list":
		fmt.Println("Preset List:")
		fmt.Println(
			`- Auth
- FileManager`,
		)
	default:
		fmt.Println("Preset not found.")
	}
}

func generateAuthPreset() {
	generatePreset(authFiles, "auth")
}

func generateFileManagerPreset() {
	generatePreset(fileManagerFiles, "filemanager")
}

func generatePreset(fsys fs.FS, packageName string) {
	writePackageName := packageName
	if u.IsPackageExist(packageName) {
		i := 1
		packageNameAlias := fmt.Sprintf("%s_%d", packageName, i)
		for u.IsPackageExist(packageNameAlias) {
			i++
			packageNameAlias = fmt.Sprintf("%s_%d", packageName, i)
		}
		writePackageName = packageNameAlias
	}
	presets, _ := fs.ReadDir(fsys, fmt.Sprintf("template/preset/%s", packageName))
	for _, file := range presets {
		fileName := file.Name()
		fileDataBuff, err := fs.ReadFile(fsys, fmt.Sprintf("template/preset/%s/%s", packageName, fileName))
		if err != nil {
			panic(err)
		}
		fileData := string(fileDataBuff)
		fileNameTranslate := strings.Replace(fileName, ".txt", "", 1)
		u.WriteFile(writePackageName, fileNameTranslate, fileData)
	}
}
