package generator_utils

import (
	"fmt"
	"os"
)

func IsPackageExist(folderName string) bool {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteFile(folderName, fileName, content string) {
	if len(folderName) > 0 {
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
	} else {
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err2 := f.WriteString(content)
		if err2 != nil {
			panic(err2)
		}
		fmt.Printf("Writing %s\n", fileName)
	}
}

func ReadFile(folderName, fileName string) string {
	if len(folderName) > 0 {
		fileDataBuff, errf := os.ReadFile(fmt.Sprintf("%s/%s", folderName, fileName))
		if errf != nil {
			panic(errf)
		}
		return string(fileDataBuff)
	} else {
		fileDataBuff, errf := os.ReadFile(fileName)
		if errf != nil {
			panic(errf)
		}
		return string(fileDataBuff)
	}
}
