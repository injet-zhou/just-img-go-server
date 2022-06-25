package dao

import (
	"errors"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"gorm.io/gorm"
)

func GetUserById(id uint) (*entity.User, error) {
	user := new(entity.User)
	if err := global.DBEngine.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DefaultUserGroup(db *gorm.DB) (*entity.UserGroup, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	var userGroup *entity.UserGroup
	if err := db.FirstOrCreate(&userGroup, &entity.UserGroup{Name: "默认", IsDefault: true}).Error; err != nil {
		return nil, err
	}
	return userGroup, nil
}
