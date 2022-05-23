package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMysqlCfg(t *testing.T) {
	initConfig("./empty.ini")
	mysqlCfg := GetMysqlCfg()
	assert.Equal(t, new(MysqlCfg), mysqlCfg, "mysqlCfg should be empty")
}

func TestGetMysqlCfg2(t *testing.T) {
	initConfig("./example.ini")
	mysqlCfg := GetMysqlCfg()
	assert.Equal(t, "foo", mysqlCfg.User, "mysqlCfg.User should be foo")
	assert.Equal(t, "bar", mysqlCfg.Password, "mysqlCfg.Password should be bar")
	assert.Equal(t, "127.0.0.1", mysqlCfg.Host, "mysqlCfg.Host should be 127.0.0.1")
	assert.Equal(t, 3306, mysqlCfg.Port, "mysqlCfg.Port should be 3306")
	assert.Equal(t, "db", mysqlCfg.Database, "mysqlCfg.Database should be db")
}
