package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/tool"
)

var client *oss.Client

func init() {
	client, _ = DefaultClient()
}

func DefaultClient() (*oss.Client, error) {
	cfg := config.GetOSSCfg()
	if cfg == nil {
		return nil, fmt.Errorf("aliyun config is nil")
	}
	if tool.IsStructEmpty(cfg) {
		return nil, fmt.Errorf("aliyun config is empty")
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

// DefaultBucket 获取默认配置bucket
func DefaultBucket() (*oss.Bucket, error) {
	cfg := config.GetOSSCfg()
	if cfg == nil {
		return nil, fmt.Errorf("aliyun config is nil")
	}
	return Bucket(cfg.BucketName)
}

func Bucket(bucketName string) (*oss.Bucket, error) {
	if client == nil {
		return nil, fmt.Errorf("aliyun client is nil")
	}
	return client.Bucket(bucketName)
}

type OSS struct {
	Client *oss.Client
}

func (o *OSS) Upload(ctx *gin.Context) (string, error) {
	f, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}
	if client == nil {
		_, err = DefaultClient()
		if err != nil {
			return "", err
		}
	}
	bucket, BucketErr := DefaultBucket()
	if BucketErr != nil {
		return "", BucketErr
	}
	file, openErr := f.Open()
	if openErr != nil {
		return "", openErr
	}
	err = bucket.PutObject(f.Filename, file)
	if err != nil {
		return "", err
	}
	return "", nil
}
