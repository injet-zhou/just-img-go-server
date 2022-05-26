package controller

import "github.com/gin-gonic/gin"

func UploadController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
