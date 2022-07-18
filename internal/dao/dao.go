package dao

import (
	"github.com/injet-zhou/just-img-go-server/global"
	"gorm.io/gorm"
)

type Dao struct {
	Engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{
		Engine: engine,
	}
}

func Default() *Dao {
	return &Dao{
		Engine: global.DBEngine,
	}
}
