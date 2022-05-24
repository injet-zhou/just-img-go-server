package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	GroupId  uint   `gorm:"not null"`
}

type UserGroup struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	IsAdmin bool
}
