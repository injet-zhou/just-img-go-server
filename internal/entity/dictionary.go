package entity

import "gorm.io/gorm"

type Dict struct {
	gorm.Model
	Key    string `gorm:"type:varchar(50);not null"`
	Value  string `gorm:"type:varchar(255);not null"`
	Type   int    `gorm:"type:int"`
	Module int    `gorm:"type:int"`
}
