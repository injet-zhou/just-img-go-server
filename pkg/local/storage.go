package local

import (
	"bytes"
	"github.com/injet-zhou/just-img-go-server/pkg"
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

func (s *Storage) Upload(file *pkg.File) (string, error) {
	root := tool.GetProjectAbsPath()
	filePath := root + "/static/" + file.Name
	data, readErr := StreamToByte(*file.File)
	if readErr != nil {
		return "", readErr
	}
	writeErr := os.WriteFile(filePath, data, 0644)
	if writeErr != nil {
		return "", writeErr
	}
	return file.Name, nil
}
