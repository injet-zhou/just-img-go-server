package dao

import (
	"errors"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"gorm.io/gorm"
)

type TokenUser struct {
	Username  string
	Password  string
	Email     string
	GroupId   uint
	Avatar    string
	Nickname  string
	UID       string
	GroupName string
	IsAdmin   bool
	IsDefault bool
}

func GetUserById(id uint) (*entity.User, error) {
	user := new(entity.User)
	if err := global.DBEngine.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(id uint) (*TokenUser, error) {
	user := new(TokenUser)
	db := global.DBEngine.Model(&entity.User{})
	db = db.Select("user.username,user.password,user.email,user.group_id,user.avatar,user.nickname,user.uid,user_group.name,user_group.is_admin,user_group.is_default")
	db = db.Joins("left join user_group on user.group_id = user_group.id")
	err := db.Find(user).Error
	if err != nil {
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
