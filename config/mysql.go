package config

import "gopkg.in/ini.v1"

// MysqlCfg mysql配置
type MysqlCfg struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// initMysqlConfig 获取mysql配置
func initMysqlConfig(cfg *ini.File) (*MysqlCfg, error) {
	mysqlCfg = new(MysqlCfg)
	mapErr := cfg.Section(mysqlSection).MapTo(mysqlCfg)
	if mapErr != nil {
		return nil, mapErr
	}
	return mysqlCfg, nil
}

func GetMysqlCfg() *MysqlCfg {
	return mysqlCfg
}
