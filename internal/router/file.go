package router

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/controller"
)

func fileRouter(r *gin.RouterGroup) {
	v1 := r.Group("/file")
	{
		v1.POST("/upload", controller.UploadController)
	}
}
