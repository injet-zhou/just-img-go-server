package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/pkg/oss"
)

func UploadController(ctx *gin.Context) {
	f, err := ctx.FormFile("file")
	if err != nil {
		ErrorResponse(ctx, 400, fmt.Errorf("upload file error: %s", err.Error()).Error())
		return
	}
	bucket, BucketErr := oss.DefaultBucket()
	if BucketErr != nil {
		ErrorResponse(ctx, 500, BucketErr.Error())
		return
	}
	file, openErr := f.Open()
	if openErr != nil {
		ErrorResponse(ctx, 500, openErr.Error())
		return
	}
	err = bucket.PutObject(f.Filename, file)
	if err != nil {
		ErrorResponse(ctx, 500, err.Error())
		return
	}
	Success(ctx, "upload success", nil)
}
