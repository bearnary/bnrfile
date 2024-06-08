package bnrfile

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/rs/xid"
	"golang.org/x/image/webp"
)

func (c *defaultClient) StoreImageBase64Local(data string) (*FileData, error) {

	b64data := data[strings.IndexByte(data, ',')+1:]

	it := PNG
	if strings.Contains(data, "image/jpeg") {
		it = JPEG
	} else if strings.Contains(data, "image/webp") {
		it = WEBP
	}

	unbased, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, fmt.Errorf("error, storage uplaod file fail: %v", err.Error())
	}

	r := bytes.NewReader(unbased)

	var im image.Image
	switch it {
	case JPEG:
		im, err = jpeg.Decode(r)
	case PNG:
		im, err = png.Decode(r)
	case WEBP:
		im, err = webp.Decode(r)
	}
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, fmt.Errorf("error, storage uplaod file fail: %v", err.Error())
	}

	tempRootDir := c.rootPath + "/" + xid.New().String()

	err = CreateDirectory(tempRootDir)
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, err
	}

	fileName := "capture.jpeg"

	dstPath := tempRootDir + "/" + fileName

	f, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, fmt.Errorf("error, storage uplaod file fail: %v", err.Error())
	}

	err = jpeg.Encode(f, im, &jpeg.Options{
		Quality: 100,
	})

	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, fmt.Errorf("error, storage uplaod file fail: %v", err.Error())
	}

	fi, err := f.Stat()
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, fmt.Errorf("error, storage uplaod file fail: %v", err.Error())
	}

	contentType := "image/jpeg"

	m := FileData{
		Name:        fileName,
		Path:        tempRootDir,
		FullPath:    dstPath,
		Size:        fi.Size(),
		ContentType: contentType,
	}

	return &m, nil
}
