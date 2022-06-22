package entity

import (
	"errors"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	if db != nil {
		return db.AutoMigrate(
			&User{},
			&UserGroup{},
			&Image{},
		)
	}
	return errors.New("db is nil")
}
