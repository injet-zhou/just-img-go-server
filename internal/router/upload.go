package router

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/controller"
)

func uploadRouter(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.POST("/upload", controller.UploadController)
	}
}
