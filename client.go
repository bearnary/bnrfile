package bnrfile

import (
	"mime/multipart"
)

type Client interface {
	StoreFileLocal(file *multipart.FileHeader) (*FileData, error)
	ScrapeImageLocal(imageUrl string, fileName string) (*FileData, error)

	StoreImageBase64Local(data string) (*FileData, error)
	ForceDeleteDirectory(name string) error
	SetMaxFileSize(size int64)
}
type defaultClient struct {
	rootPath    string
	maxFileSize int64
}

func NewClient(rootPath string) (Client, error) {

	err := CreateDirectory(rootPath)
	if err != nil {
		return nil, err
	}
	return &defaultClient{
		rootPath: rootPath,
	}, nil
}

func (c *defaultClient) SetMaxFileSize(size int64) {
	c.maxFileSize = size
}
