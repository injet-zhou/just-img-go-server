package main

import (
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/internal/router"
	"log"
)

func main() {
	r := router.RouteSetup()
	err := r.Run(":" + config.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
