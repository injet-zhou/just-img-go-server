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
	mysqlCfg   *MysqlCfg
	ossCfg     *OSSCfg
	cosCfg     *COSCfg
	qiniuCfg   *QiniuCfg
	upyunCfg   *UpyunCfg
	redisCfg   *RedisCfg
	jwtCfg     *JwtCfg
	accountCfg *AccountCfg
)

type JwtCfg struct {
	Secret string
	Expire int64
}

type AccountCfg struct {
	Username string
	Password string
}

func initAccountCfg(cfg *ini.File) (*AccountCfg, error) {
	accountCfg = new(AccountCfg)
	mapErr := cfg.Section("account").MapTo(accountCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return accountCfg, nil
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
	var ossErr error
	ossCfg, ossErr = initOSSCfg(cfg)
	if ossErr != nil {
		warn(ossSection, err)
	}
	var cosErr error
	cosCfg, cosErr = initCOSCfg(cfg)
	if cosErr != nil {
		warn(cosSection, err)
	}
	var qiniuErr error
	qiniuCfg, qiniuErr = initQiniuCfg(cfg)
	if qiniuErr != nil {
		warn(qiniuSection, err)
	}
	var upyunErr error
	upyunCfg, upyunErr = initUpyunCfg(cfg)
	if upyunErr != nil {
		warn(qiniuSection, err)
	}
	var redisErr error
	redisCfg, redisErr = initRedisConfig(cfg)
	if redisErr != nil {
		warn("redis", err)
	}
	var jwtErr error
	jwtCfg, jwtErr = initJwtCfg(cfg)
	if jwtErr != nil {
		warn("jwt", err)
	}
	var accountErr error
	accountCfg, err = initAccountCfg(cfg)
	if accountErr != nil {
		warn("account", err)
	}
}

func GetJwtCfg() *JwtCfg {
	return jwtCfg
}

func GetAccountCfg() *AccountCfg {
	return accountCfg
}

func init() {
	initConfig("")
}
