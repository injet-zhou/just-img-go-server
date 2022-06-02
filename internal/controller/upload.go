package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/pkg/upload"
)

func UploadController(ctx *gin.Context) {
	platformType, ok := ctx.GetPostForm("platform")
	if !ok {
		ErrorResponse(ctx, 400, "platform is required")
		return
	}
	num, parseErr := fmt.Sscanf(platformType, "%d", &platformType)
	if num <= 0 || parseErr != nil {
		ErrorResponse(ctx, 400, "platform is invalid")
		return
	}
	uploader := upload.NewUploader(config.Local)
	res, uploadErr := uploader.Upload(ctx)
	if uploadErr != nil {
		ErrorResponse(ctx, 500, uploadErr.Error())
		return
	}
	Success(ctx, "upload success", res)
}
