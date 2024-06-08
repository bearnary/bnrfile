package bnrfile

import (
	"os"
)

func CreateDirectory(path string) error {

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	//Create a folder/directory at a full qualified path
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
