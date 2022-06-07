package main

import (
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/global"
	"github.com/injet-zhou/just-img-go-server/internal/router"
	"log"
)

func init() {
	err := global.DBSetup()
	if err != nil {
		log.Fatal("db setup error: ", err)
	}

}

func main() {
	r := router.RouteSetup()
	err := r.Run(":" + config.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
