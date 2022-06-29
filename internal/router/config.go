package router

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/controller"
)

func ConfigRouter(r *gin.RouterGroup) {
	cfgRoute := r.Group("/config")
	{
		cfgRoute.GET("/platforms", controller.PlatformsController)
	}
}
