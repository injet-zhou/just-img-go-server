package upyun

import (
	"fmt"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/pkg"
	"github.com/injet-zhou/just-img-go-server/tool"
	"github.com/upyun/go-sdk/v3/upyun"
	"strings"
)

var client *upyun.UpYun

func DefaultClient() (*upyun.UpYun, error) {
	if client == nil {
		cfg := config.GetUpyunCfg()
		if cfg == nil {
			return nil, fmt.Errorf("upyun config is nil")
		}
		if tool.IsStructEmpty(cfg) {
			return nil, fmt.Errorf("upyun config is empty")
		}
		up, err := NewClient(cfg)
		if err != nil {
			return nil, err
		}
		client = up
	}
	return client, nil
}

func NewClient(cfg *config.UpyunCfg) (*upyun.UpYun, error) {
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   cfg.Bucket,
		Operator: cfg.Operator,
		Password: cfg.Password,
	})
	return up, nil
}

type Upyun struct {
}

func (u *Upyun) Upload(file *pkg.File) (string, error) {
	up, err := DefaultClient()
	if err != nil {
		return "", err
	}
	uploadErr := up.Put(&upyun.PutObjectConfig{
		Path:   "/" + file.Name,
		Reader: *file.File,
	})
	if uploadErr != nil {
		return "", uploadErr
	}
	cfg := config.GetUpyunCfg()
	url := ""
	if cfg.BaseURL != "" {
		if strings.HasSuffix(cfg.BaseURL, "/") {
			url = cfg.BaseURL + file.Name
		} else {
			url = cfg.BaseURL + "/" + file.Name
		}
	}
	return url, nil
}
