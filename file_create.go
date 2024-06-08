package bnrfile

import (
	"os"
)

func CreateFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()
	return nil
}
