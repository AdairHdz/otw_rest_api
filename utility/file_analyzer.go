package utility

import (
	"io"
	"os"
)

func DirIsEmpty(fileName string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}

	defer file.Close()

	_, err = file.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func IsImage(fileExtension string) bool {
	if fileExtension == ".png" || fileExtension == ".jpeg" || fileExtension == ".jpg" {
		return true
	}

	return false
}