package service

import "github.com/injet-zhou/just-img-go-server/config"

type Platform struct {
	Label string              `json:"label"`
	Type  config.PlatformType `json:"type"`
}

type GetCfg func()

func GetPlatforms() []Platform {
	return []Platform{
		{Label: "阿里云", Type: config.OSS},
		{Label: "腾讯云", Type: config.COS},
		{Label: "七牛云", Type: config.QINIU},
		{Label: "又拍云", Type: config.UPYUN},
		{Label: "本地", Type: config.Local},
	}
}
