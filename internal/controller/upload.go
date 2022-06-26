package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/pkg"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"github.com/injet-zhou/just-img-go-server/pkg/upload"
	"go.uber.org/zap"
)

func UploadController(ctx *gin.Context) {
	platformType, ok := ctx.GetPostForm("platform")
	if !ok {
		ErrorResponse(ctx, 400, "platform is required")
		return
	}
	num, parseErr := fmt.Sscanf(platformType, "%d", &platformType)
	if num <= 0 || parseErr != nil {
		logger.Error("platform is invalid", zap.String("platform", platformType))
		ErrorResponse(ctx, 400, "platform is invalid")
		return
	}
	uploader := upload.NewUploader(config.PlatformType(num))
	file, err := pkg.GetFile(ctx)
	if err != nil {
		logger.Error("get file error", zap.String("err", err.Error()))
		ErrorResponse(ctx, 400, err.Error())
		return
	}
	res, uploadErr := uploader.Upload(file)
	if uploadErr != nil {
		ErrorResponse(ctx, 500, uploadErr.Error())
		return
	}
	Success(ctx, "upload success", res)
}
