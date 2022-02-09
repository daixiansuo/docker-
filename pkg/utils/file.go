package utils

import (
	"fmt"
	"os"
	"path"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PathInit(filePath string) error {

	if !Exists(filePath) {

		// split dir & file
		dir, _ := path.Split(filePath)

		// create dir
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("mkdir path: %s, error %v", dir, err)
		}
		// create file
		if _, err := os.Create(filePath); err != nil {
			return fmt.Errorf("create file %s error %v", filePath, err)
		}
	}

	return nil
}
