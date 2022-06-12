package config

import "gopkg.in/ini.v1"

// RedisCfg redis配置
type RedisCfg struct {
	Host     string
	Port     int
	Password string
	Username string
	DB       int
}

func initRedisConfig(cfg *ini.File) (*RedisCfg, error) {
	redisCfg = new(RedisCfg)
	mapErr := cfg.Section(redisSection).MapTo(redisCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return redisCfg, nil
}

func GetRedisCfg() *RedisCfg {
	return redisCfg
}
