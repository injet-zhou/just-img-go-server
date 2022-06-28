package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"mime/multipart"
	"strings"
	"time"
)

type File struct {
	File         *multipart.File
	Name         string
	Size         int64
	Type         string
	URL          string
	Path         string
	OriginalName string
}

func path() string {
	now := time.Now()
	return now.Format("2006/01/02/")
}

func newFilename(filename string) string {
	strs := strings.Split(filename, ".")
	if len(strs) > 1 {
		guid := xid.New()
		return guid.String() + "." + strs[len(strs)-1]
	}
	return filename
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
		File:         &file,
		OriginalName: filename,
		Size:         f.Size,
		Type:         f.Header.Get("Data-Type"),
		Path:         path(),
		Name:         newFilename(filename),
	}, nil
}
