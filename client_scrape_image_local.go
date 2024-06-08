package bnrfile

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/xid"
)

func (c *defaultClient) ScrapeImageLocal(imageUrl string, fileName string) (*FileData, error) {

	response, err := http.Get(imageUrl)
	if err != nil {
		return nil, fmt.Errorf("error, image cannot be loaded: %v", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("error, image cannot be loaded")
	}

	tempRootDir := c.rootPath + "/" + xid.New().String()

	err = CreateDirectory(tempRootDir)
	if err != nil {
		return nil, err
	}

	// Destination
	dstPath := tempRootDir + "/" + fileName + ".png"
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, fmt.Errorf("error, file not found: %v", err.Error())
	}
	defer dst.Close()

	// Copy
	_, err = io.Copy(dst, response.Body)
	if err != nil {
		return nil, fmt.Errorf("error, image cannot be loaded: %v", err.Error())
	}

	fileInfo, err := dst.Stat()
	if err != nil {
		return nil, fmt.Errorf("error, image cannot be loaded: %v", err.Error())
	}

	m := FileData{
		Name:        fileName,
		Path:        tempRootDir,
		FullPath:    dstPath,
		Size:        fileInfo.Size(),
		ContentType: "png",
	}

	return &m, nil
}
