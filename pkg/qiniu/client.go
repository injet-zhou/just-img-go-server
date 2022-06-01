package qiniu

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/tool"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
}

func (q *Qiniu) Upload(ctx *gin.Context) (string, error) {
	cfg := config.GetQiniuCfg()
	if cfg == nil {
		return "", fmt.Errorf("qiniu config is not set")
	}
	if tool.IsStructEmpty(cfg) {
		return "", fmt.Errorf("qiniu config is empty")
	}
	putPolicy := storage.PutPolicy{
		Scope: cfg.Bucket,
	}
	mac := qbox.NewMac(cfg.AccessKey, cfg.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	conf := storage.Config{}
	// 空间对应的机房
	conf.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	conf.UseHTTPS = true
	// 上传是否使用CDN上传加速
	conf.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&conf)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	f, openErr := file.Open()
	if openErr != nil {
		return "", openErr
	}
	data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	uploadErr := formUploader.Put(context.Background(), &ret, upToken, fileName, f, dataLen, &putExtra)
	if err != nil {
		fmt.Println(uploadErr)
		return "", uploadErr
	}
	return ret.Key, nil
}
