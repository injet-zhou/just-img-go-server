package entity

import (
	"fmt"
	"github.com/injet-zhou/just-img-go-server/tool"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	GroupId  uint   `gorm:"not null"`
	Avatar   string `gorm:"type:varchar(255)"`
	Nickname string `gorm:"type:varchar(100)"`
	UID      string `gorm:"type:varchar(100)"`
}

type SafeUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	UserId   uint   `json:"userId"`
	GroupId  uint   `json:"groupId"`
	Token    string `json:"token"`
}

type UserGroup struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	IsAdmin bool
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	guid := xid.New()
	u.UID = guid.String()
	u.Password = tool.MD5(u.Password + u.UID)
	return nil
}

func (u *User) SafeInfo() *SafeUser {
	return &SafeUser{
		Username: u.Username,
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		UserId:   u.ID,
		GroupId:  u.GroupId,
	}
}

func (u *User) GetByUsername(db *gorm.DB) (*User, error) {
	if strings.Trim(u.Username, " ") == "" {
		return nil, fmt.Errorf("username is required")
	}
	if db == nil {
		return nil, fmt.Errorf("db is required")
	}
	var user User
	err := db.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) GetByLoginName(db *gorm.DB) (*User, error) {
	if db == nil {
		return nil, fmt.Errorf("db is required")
	}
	var user User
	err := db.Where("username = ? or email = ?", u.Username, u.Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Create(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("db is required")
	}
	return db.Create(u).Error
}
