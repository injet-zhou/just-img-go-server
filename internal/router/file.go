package router

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/controller"
	"github.com/injet-zhou/just-img-go-server/internal/middleware"
)

func fileRouter(r *gin.RouterGroup) {
	v1 := r.Group("/file")
	{
		v1.Use(middleware.AuthMiddleware())
		v1.POST("/upload", controller.UploadController)
		v1.POST("/image/list", controller.ImageListController)
	}
}
