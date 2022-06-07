package entity

import "gorm.io/gorm"

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

func (u *User) GetByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
