package dao

import (
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
)

func GetUserById(id uint) (*entity.User, error) {
	user := new(entity.User)
	if err := global.DBEngine.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
