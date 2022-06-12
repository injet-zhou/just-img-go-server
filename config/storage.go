package config

import "gopkg.in/ini.v1"

// OSSCfg 阿里云OSS配置
type OSSCfg struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	AccessDomain    string
}

// COSCfg 腾讯云COS配置
type COSCfg struct {
	Region    string
	SecretId  string
	SecretKey string
	Bucket    string
	BaseURL   string
}

// QiniuCfg 七牛云配置
type QiniuCfg struct {
	AccessKey    string
	SecretKey    string
	Bucket       string
	AccessDomain string
}

// UpyunCfg 又拍云配置
type UpyunCfg struct {
	Bucket   string
	Operator string
	Password string
}

func initOSSCfg(cfg *ini.File) (*OSSCfg, error) {
	ossCfg = new(OSSCfg)
	mapErr := cfg.Section(ossSection).MapTo(ossCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return ossCfg, nil
}

func initCOSCfg(cfg *ini.File) (*COSCfg, error) {
	cosCfg = new(COSCfg)
	mapErr := cfg.Section(cosSection).MapTo(cosCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return cosCfg, nil
}

func initQiniuCfg(cfg *ini.File) (*QiniuCfg, error) {
	qiniuCfg = new(QiniuCfg)
	mapErr := cfg.Section(qiniuSection).MapTo(qiniuCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return qiniuCfg, nil
}

func initUpyunCfg(cfg *ini.File) (*UpyunCfg, error) {
	upyunCfg = new(UpyunCfg)
	mapErr := cfg.Section(upyunSection).MapTo(upyunCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return upyunCfg, nil
}

func GetOSSCfg() *OSSCfg {
	return ossCfg
}

func GetCOSCfg() *COSCfg {
	return cosCfg
}

func GetQiniuCfg() *QiniuCfg {
	return qiniuCfg
}

func GetUpyunCfg() *UpyunCfg {
	return upyunCfg
}
