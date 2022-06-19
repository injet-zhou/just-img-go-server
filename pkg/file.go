package pkg

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type File struct {
	File *multipart.File
	Name string
	Size int64
	Type string
	URL  string
	Path string
}

func GetFile(ctx *gin.Context) (*File, error) {
	f, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}
	filename := f.Filename
	file, openErr := f.Open()
	if openErr != nil {
		return nil, openErr
	}
	return &File{
		File: &file,
		Name: filename,
		Size: f.Size,
		Type: f.Header.Get("Content-Type"),
	}, nil
}
