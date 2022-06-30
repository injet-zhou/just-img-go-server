package qiniu

import (
	"context"
	"fmt"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/pkg"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"github.com/injet-zhou/just-img-go-server/tool"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"strings"
)

type Qiniu struct {
}

func (q *Qiniu) Upload(file *pkg.File) (string, error) {
	log := logger.Default()
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
	var err error
	filename := file.Path + file.Name
	uploadErr := formUploader.Put(context.Background(), &ret, upToken, filename, *file.File, file.Size, &putExtra)
	if err != nil {
		log.Error("upload file error:", zap.String("err", err.Error()))
		return "", uploadErr
	}
	url := ""
	if cfg.BaseURL != "" {
		if strings.HasSuffix(cfg.BaseURL, "/") {
			url = cfg.BaseURL + filename
		} else {
			url = cfg.BaseURL + "/" + filename
		}
	}
	return url, nil
}
