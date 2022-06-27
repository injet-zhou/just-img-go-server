package service

import (
	"errors"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/pkg"
)

type UploadInfo struct {
	File *pkg.File
	User *entity.User
	IP   string
}

func SaveUploadInfo(uploadInfo *UploadInfo) error {
	if uploadInfo == nil {
		return errors.New("uploadInfo is nil")
	}
	if uploadInfo.File == nil {
		return errors.New("uploadInfo.File is nil")
	}
	if uploadInfo.User == nil {
		return errors.New("uploadInfo.User is nil")
	}
	img := &entity.Image{
		Name:         uploadInfo.File.Name,
		UserId:       uploadInfo.User.ID,
		Size:         uploadInfo.File.Size,
		MimeType:     uploadInfo.File.Type,
		Path:         uploadInfo.File.Path,
		URL:          uploadInfo.File.URL,
		UploadIP:     uploadInfo.IP,
		OriginalName: uploadInfo.File.OriginalName,
	}
	_, err := img.Create(global.DBEngine)
	if err != nil {
		return err
	}
	return nil
}
