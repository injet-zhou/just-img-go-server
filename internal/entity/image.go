package entity

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserId       uint   `gorm:"not null"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Path         string `gorm:"type:varchar(255);not null" json:"path"`
	Size         int64  `gorm:"type:bigint" json:"size"`
	OriginalName string `gorm:"type:varchar(255)" json:"originalName"`
	MimeType     string `gorm:"type:varchar(100);not null" json:"mimeType"`
	URL          string `gorm:"type:varchar(255);not null" json:"URL"`
	MD5          string `gorm:"type:varchar(100);not null" json:"MD5"`
	UploadIP     string `gorm:"type:varchar(32)" json:"uploadIP"`
	GroupId      uint   `gorm:"not null" json:"groupId"`
}

func (i *Image) Create(db *gorm.DB) (*Image, error) {
	err := db.Create(&i).Error
	return i, err
}
