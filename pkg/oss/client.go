package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/injet-zhou/just-img-go-server/config"
)

var client *oss.Client

func init() {
	client, _ = DefaultClient()
}

func DefaultClient() (*oss.Client, error) {
	cfg := config.GetOSSCfg()
	if cfg == nil {
		return nil, fmt.Errorf("oss config is nil")
	}
	return NewClient(cfg)
}

// NewClient 获取OSS client
func NewClient(ossCfg *config.OSSCfg) (*oss.Client, error) {
	endpoint := ""
	accessKeyID := ""
	accessKeySecret := ""
	if ossCfg != nil {
		endpoint = ossCfg.Endpoint
		accessKeyID = ossCfg.AccessKeyId
		accessKeySecret = ossCfg.AccessKeySecret
	}
	if client == nil {
		var err error
		client, err = oss.New(endpoint, accessKeyID, accessKeySecret)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func DefaultBucket() (*oss.Bucket, error) {
	cfg := config.GetOSSCfg()
	if cfg == nil {
		return nil, fmt.Errorf("oss config is nil")
	}
	return Bucket(cfg.BucketName)
}

func Bucket(bucketName string) (*oss.Bucket, error) {
	if client == nil {
		return nil, fmt.Errorf("oss client is nil")
	}
	return client.Bucket(bucketName)
}
