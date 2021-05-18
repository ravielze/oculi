package filemanager

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/ravielze/oculi/common/radix36"
)

func GenerateFileName(ext string) string {
	name, err := radix36.EncodeUUID4()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s-%s-%s%s", name[0:9], name[9:17], name[17:], ext)
}

func DownloadFile(URL, path string) (string, error) {
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
		panic("received non 200 response code")
	}

	contents, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return "", err2
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(contents))

	fileType, err3 := mimetype.DetectReader(response.Body)
	if err3 != nil {
		return "", err3
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(contents))

	//Create a empty file
	fileName := fmt.Sprintf("%s/%s", path, GenerateFileName(fileType.Extension()))

	file, err4 := os.Create(fileName)
	if err4 != nil {
		return "", err4
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err5 := io.Copy(file, response.Body)
	if err5 != nil {
		return "", err5
	}

	return fileName, nil
}
