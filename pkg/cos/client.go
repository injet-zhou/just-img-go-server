package cos

import (
	"context"
	"fmt"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/pkg"
	"github.com/injet-zhou/just-img-go-server/tool"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client *cos.Client

type COS struct {
}

func DefaultClient() (*cos.Client, error) {
	cfg := config.GetCOSCfg()
	if cfg == nil {
		return nil, fmt.Errorf("aliyun config is nil")
	}
	if tool.IsStructEmpty(cfg) {
		return nil, fmt.Errorf("aliyun config is empty")
	}
	return NewClient(cfg)
}

func NewClient(cfg *config.COSCfg) (*cos.Client, error) {
	if client != nil {
		return client, nil
	}
	u, _ := url.Parse(cfg.BaseURL)
	b := &cos.BaseURL{BucketURL: u}
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretId,
			SecretKey: cfg.SecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: false,
				ResponseBody:   false,
			},
		},
	})
	return client, nil
}

func (c *COS) Upload(file *pkg.File) (string, error) {
	if client == nil {
		return "", fmt.Errorf("tencent cos client is nil")
	}
	filename := file.Path + file.Name
	res, err := client.Object.Put(context.Background(), filename, *file.File, nil)
	if err != nil {
		return "", err
	}
	defer func() {
		closeErr := res.Body.Close()
		if closeErr != nil {
			fmt.Printf("close response body failed, %v\n", closeErr)
		}
	}()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}
	fmt.Printf("%s\n", body)
	URL := ""
	if cfg := config.GetCOSCfg(); cfg != nil {
		if cfg.BaseURL == "" {
			return URL, nil
		}
		URL = strings.TrimRight(cfg.BaseURL, "/") + "/" + filename
	}
	return URL, nil
}
