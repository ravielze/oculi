package generator_utils

import (
	"fmt"
	"os"
)

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

func ReadFile(folderName, fileName string) string {
	fileDataBuff, errf := os.ReadFile(fmt.Sprintf("%s/%s", folderName, fileName))
	if errf != nil {
		panic(errf)
	}
	fileData := string(fileDataBuff)
	return fileData
}