package service

import "github.com/injet-zhou/just-img-go-server/config"

type Platform struct {
	Name string              `json:"name"`
	Type config.PlatformType `json:"type"`
}

type GetCfg func()

func GetPlatforms() []Platform {
	return []Platform{
		{Name: "阿里云", Type: config.OSS},
		{Name: "腾讯云", Type: config.COS},
		{Name: "七牛云", Type: config.QINIU},
		{Name: "又拍云", Type: config.UPYUN},
		{Name: "本地", Type: config.Local},
	}
}
