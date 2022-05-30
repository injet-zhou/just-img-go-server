package local

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/tool"
	"io"
	"os"
)

type Storage struct {
}

func StreamToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *Storage) Upload(ctx *gin.Context) (string, error) {
	f, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}
	file, openErr := f.Open()
	if openErr != nil {
		return "", openErr
	}
	defer file.Close()
	fileName := f.Filename
	root := tool.GetProjectAbsPath()
	filePath := root + "/static/" + fileName
	data, readErr := StreamToByte(file)
	if readErr != nil {
		return "", readErr
	}
	writeErr := os.WriteFile(filePath, data, 0644)
	if writeErr != nil {
		return "", writeErr
	}
	return fileName, nil
}
