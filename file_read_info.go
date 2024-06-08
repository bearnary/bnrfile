package bnrfile

import (
	"fmt"
	"os"
	"strings"

	"github.com/h2non/filetype"
)

func ReadFileInfo(fullPath string) (*FileData, error) {
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get the content
	ct, err := GetFileContentType(f)

	if !filetype.IsMIMESupported(ct) {
		return nil, fmt.Errorf("error, unsupported file type: %v", ct)
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	basePath := strings.Replace(fullPath, fi.Name(), "", -1)

	return &FileData{
		Name:        fi.Name(),
		Path:        basePath,
		FullPath:    fullPath,
		Size:        fi.Size(),
		ContentType: ct,
	}, nil
}
