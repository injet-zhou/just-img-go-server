package entity

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserId       uint   `json:"user_id" gorm:"not null"`
	Name         string `gorm:"type:varchar(100);not null"`
	Path         string `gorm:"type:varchar(255);not null"`
	Size         int64  `gorm:"type:bigint"`
	OriginalName string `gorm:"type:varchar(255)"`
	MimeType     string `gorm:"type:varchar(100);not null"`
	URL          string `gorm:"type:varchar(255);not null"`
	MD5          string `gorm:"type:varchar(100);not null"`
	UploadIP     string `gorm:"type:varchar(32)"`
}

func (i *Image) Create(db *gorm.DB) (*Image, error) {
	err := db.Create(&i).Error
	return i, err
}
