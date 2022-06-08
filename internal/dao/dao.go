package dao

import (
	"github.com/injet-zhou/just-img-go-server/global"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{
		engine: engine,
	}
}

func Default() *Dao {
	return &Dao{
		engine: global.DBEngine,
	}
}
