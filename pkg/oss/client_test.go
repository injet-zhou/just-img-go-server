package oss

import (
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	cfg := &config.OSSCfg{
		Endpoint:        "test",
		AccessKeyId:     "test",
		AccessKeySecret: "test",
	}
	client, _ := NewClient(cfg)
	assert.NotNil(t, client)
}

func TestDefaultClient(t *testing.T) {
	client, _ := DefaultClient()
	assert.NotNil(t, client)
}
