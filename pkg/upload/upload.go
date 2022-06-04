package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/pkg/aliyun"
	"github.com/injet-zhou/just-img-go-server/pkg/cos"
	"github.com/injet-zhou/just-img-go-server/pkg/local"
	"github.com/injet-zhou/just-img-go-server/pkg/qiniu"
	"github.com/injet-zhou/just-img-go-server/pkg/upyun"
)

type Uploader interface {
	Upload(ctx *gin.Context) (string, error)
}

func NewUploader(platformType config.PlatformType) Uploader {
	switch platformType {
	case config.OSS:
		return &aliyun.OSS{}
	case config.COS:
		return &cos.COS{}
	case config.QINIU:
		return &qiniu.Qiniu{}
	case config.UPYUN:
		return &upyun.Upyun{}
	case config.Local:
		return &local.Storage{}
	default:
		return &local.Storage{}
	}
}
