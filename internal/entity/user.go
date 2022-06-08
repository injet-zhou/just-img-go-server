package entity

import (
	"fmt"
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
}

type UserGroup struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	IsAdmin bool
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
