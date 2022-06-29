package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/internal/service"
)

func PlatformsController(ctx *gin.Context) {
	Success(ctx, "success", service.GetPlatforms())
}
