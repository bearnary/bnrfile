package bnrfile

import (
	"os"
)

func ForceDeleteDirectory(name string) error {

	if _, err := os.Stat(name); !os.IsNotExist(err) {

		err = os.RemoveAll(name)
		if err != nil {
			return err
		}
	}

	return nil
}
