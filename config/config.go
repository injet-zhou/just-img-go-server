package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

const (
	mysqlSection      = "mysql"
	ossSection        = "oss"
	cosSection        = "cos"
	qiniuSection      = "qiniu"
	upyunSection      = "upyun"
	defaultConfigPath = "./env.ini"
)

// MysqlCfg mysql配置
type MysqlCfg struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// OSSCfg 阿里云OSS配置
type OSSCfg struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

// COSCfg 腾讯云COS配置
type COSCfg struct {
	Region    string
	AppId     string
	SecretId  string
	SecretKey string
	Bucket    string
}

// QiniuCfg 七牛云配置
type QiniuCfg struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

// UpyunCfg 又拍云配置
type UpyunCfg struct {
	Bucket   string
	Operator string
	Password string
}

var (
	mysqlCfg *MysqlCfg
	ossCfg   *OSSCfg
	cosCfg   *COSCfg
	qiniuCfg *QiniuCfg
	upyunCfg *UpyunCfg
)

// initMysqlConfig 获取mysql配置
// defaultConfigPath 配置文件路径
func initMysqlConfig(cfg *ini.File) (*MysqlCfg, error) {
	mysqlCfg = new(MysqlCfg)
	mapErr := cfg.Section(mysqlSection).MapTo(mysqlCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return mysqlCfg, nil
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

func warn(service string, err error) {
	fmt.Printf("%s config failed: %v\n", service, err)
	fmt.Println("Skipping...")
}

func initConfig(configPath string) {
	if configPath == "" {
		configPath = defaultConfigPath
	}
	cfg, err := ini.Load(configPath)
	if err != nil {
		panic(err)
	}
	mysqlCfg, err = initMysqlConfig(cfg)
	if err != nil {
		warn(mysqlSection, err)
	}
	ossCfg, err = initOSSCfg(cfg)
	if err != nil {
		warn(ossSection, err)
	}
	cosCfg, err = initCOSCfg(cfg)
	if err != nil {
		warn(cosSection, err)
	}
	qiniuCfg, err = initQiniuCfg(cfg)
	if err != nil {
		warn(qiniuSection, err)
	}
	upyunCfg, err = initUpyunCfg(cfg)
	if err != nil {
		warn(qiniuSection, err)
	}
}

func GetMysqlCfg() *MysqlCfg {
	return mysqlCfg
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

func init() {
	initConfig("")
}
