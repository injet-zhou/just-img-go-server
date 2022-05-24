package qiniu

import (
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
)

var client *qbox.Mac

func init() {
	accessKey := ""
	secretKey := ""
	qiniuCfg := config.GetQiniuCfg()
	if qiniuCfg != nil {
		accessKey = qiniuCfg.AccessKey
		secretKey = qiniuCfg.SecretKey
		client = qbox.NewMac(accessKey, secretKey)
	}
}
