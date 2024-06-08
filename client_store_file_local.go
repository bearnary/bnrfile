package bnrfile

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/rs/xid"
)

func (c *defaultClient) StoreFileLocal(file *multipart.FileHeader) (*FileData, error) {

	if c.maxFileSize > 0 && file.Size > c.maxFileSize {
		return nil, errors.New("file exceed size limit")
	}
	tempRootDir := c.rootPath + "/" + xid.New().String()

	err := CreateDirectory(tempRootDir)
	if err != nil {
		return nil, err
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("error, storage uplaod file fail: %v", err.Error())
	}
	defer src.Close()

	contentType := file.Header.Get("Content-Type")

	// Destination
	dstPath := tempRootDir + "/" + file.Filename
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, fmt.Errorf("error, file not found: %v", err.Error())
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("error, storage upload file fail: %v", err.Error())
	}

	m := FileData{
		Name:        file.Filename,
		Path:        tempRootDir,
		FullPath:    dstPath,
		Size:        file.Size,
		ContentType: contentType,
	}

	return &m, nil
}
