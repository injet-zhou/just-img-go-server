package service

import (
	"errors"
	"github.com/injet-zhou/just-img-go-server/global"
	DAO "github.com/injet-zhou/just-img-go-server/internal/dao"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/pkg"
)

type ImagesRequest struct {
	Username     string `json:"username"`
	GroupName    string `json:"groupName"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
	OriginalName string `json:"originalName"`
	UploadIP     string `json:"uploadIP"`
}

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
		GroupId:      uploadInfo.User.GroupId,
	}
	_, err := img.Create(global.DBEngine)
	if err != nil {
		return err
	}
	return nil
}

func ImageList(req *ImagesRequest) ([]*DAO.ImagesResponse, error) {
	var images []*DAO.ImagesResponse
	paginate := &pkg.Pagination{
		Limit: req.Limit,
		Page:  req.Page,
	}
	db := global.DBEngine.Model(&entity.Image{})
	db = db.Select("image.id, image.name, image.path, image.size, image.original_name, image.mime_type, image.url, image.md5, image.upload_ip, user.username, user_group.name as group_name")
	db = db.Joins("left join user on user.id = image.user_id")
	db = db.Joins("left join user_group on user_group.id = image.group_id")
	if req.Username != "" {
		db = db.Where("user.username = ?", req.Username)
	}
	if req.GroupName != "" {
		db = db.Where("user_group.name = ?", req.GroupName)
	}
	if req.OriginalName != "" {
		db = db.Where("image.original_name = ?", req.OriginalName)
	}
	if req.UploadIP != "" {
		db = db.Where("image.upload_ip = ?", req.UploadIP)
	}
	dao := &DAO.Dao{
		Engine: db,
	}
	var err error
	images, err = dao.ImageList(paginate)
	if err != nil {
		return nil, err
	}
	return images, nil
}
