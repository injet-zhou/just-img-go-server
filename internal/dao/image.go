package dao

import (
	"github.com/injet-zhou/just-img-go-server/pkg"
)

type ImagesResponse struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"userId"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Size         int64  `json:"size"`
	OriginalName string `json:"originalName"`
	MimeType     string `json:"mimeType"`
	URL          string `json:"url"`
	MD5          string `json:"MD5"`
	UploadIP     string `json:"uploadIP"`
	GroupId      uint   `json:"groupId"`
	GroupName    string `json:"groupName"`
	Username     string `json:"username"`
}

func (d *Dao) ImageList(paginate *pkg.Pagination) ([]*ImagesResponse, error) {
	var images []*ImagesResponse
	images = make([]*ImagesResponse, paginate.GetLimit())
	err := d.Engine.Scopes(pkg.Paginate(paginate, d.Engine)).Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}
