package helpers

import (
	"errors"
	"io/ioutil"
	"os"
)

func WriteJsonFile(fileType string, fileName string, isTruncate bool, content string) {
	filePath := "./files/" + fileName
	files, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	PanicIfError(err)
	defer files.Close()

	files.WriteString(content)
}

func ReadJsonFile(fileName string) []byte {
	files, err := ioutil.ReadFile(fileName)
	PanicIfError(err)

	return files
}

func CheckDirOrFileExists(fileName string) bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
