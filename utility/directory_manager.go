package utility

import "os"

func CreateDirectory(path string) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		directoryCreationError := os.Mkdir(path, 777)
		if directoryCreationError != nil {
			return directoryCreationError
		}
		return nil
	}
	return err 
}