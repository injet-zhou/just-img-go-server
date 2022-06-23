package main

import (
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/internal/router"
	"log"
)

func init() {
	err := global.DBSetup()
	if err != nil {
		log.Fatal("db setup error: ", err)
	}
	setupTblErr := entity.InitTables(global.DBEngine)
	if setupTblErr != nil {
		log.Fatal("init tables error: ", setupTblErr)
	}
	createAdminUserErr := entity.CreateAdminUser()
	if createAdminUserErr != nil {
		log.Fatal("create admin user error: ", createAdminUserErr)
	}
}

func main() {
	r := router.RouteSetup()
	err := r.Run(":" + config.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
