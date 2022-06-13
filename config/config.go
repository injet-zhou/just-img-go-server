package config

import (
	"fmt"
	"github.com/injet-zhou/just-img-go-server/tool"
	"gopkg.in/ini.v1"
)

const (
	mysqlSection = "mysql"
	ossSection   = "oss"
	cosSection   = "cos"
	qiniuSection = "qiniu"
	upyunSection = "upyun"
	redisSection = "redis"
)

var (
	mysqlCfg *MysqlCfg
	ossCfg   *OSSCfg
	cosCfg   *COSCfg
	qiniuCfg *QiniuCfg
	upyunCfg *UpyunCfg
	redisCfg *RedisCfg
	jwtCfg   *JwtCfg
)

type JwtCfg struct {
	Secret string
	Expire int64
}

func initJwtCfg(cfg *ini.File) (*JwtCfg, error) {
	jwtCfg = new(JwtCfg)
	mapErr := cfg.Section("jwt").MapTo(jwtCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return jwtCfg, nil
}

// defaultConfigPath 配置文件路径
func defaultConfigPath() string {
	root := tool.GetProjectAbsPath()
	return root + "/config/env.ini"
}

func warn(service string, err error) {
	fmt.Printf("%s config failed: %v\n", service, err)
	fmt.Println("Skipping...")
}

func initConfig(configPath string) {
	if configPath == "" {
		configPath = defaultConfigPath()
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
	redisCfg, err = initRedisConfig(cfg)
	if err != nil {
		warn("redis", err)
	}
	jwtCfg, err = initJwtCfg(cfg)
	if err != nil {
		warn("jwt", err)
	}
}

func GetJwtCfg() *JwtCfg {
	return jwtCfg
}

func init() {
	initConfig("")
}
