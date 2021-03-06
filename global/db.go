package global

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/tool"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBEngine *gorm.DB
var RedisEngine *redis.Client

func DefaultDB() (*gorm.DB, error) {
	if DBEngine != nil {
		return DBEngine, nil
	}
	cfg := config.GetMysqlCfg()
	if cfg == nil {
		return nil, fmt.Errorf("mysql config is nil")
	}
	if tool.IsStructEmpty(cfg) {
		return nil, fmt.Errorf("mysql config is empty")
	}
	db, err := NewDbEngine(cfg)
	if err != nil {
		return nil, err
	}
	DBEngine = db
	return db, nil
}

func DefaultRedis() (*redis.Client, error) {
	if RedisEngine != nil {
		return RedisEngine, nil
	}
	cfg := config.GetRedisCfg()
	if cfg == nil {
		return nil, fmt.Errorf("redis config is nil")
	}
	redisClient, err := NewRedisEngine(cfg)
	if err != nil {
		return nil, err
	}
	RedisEngine = redisClient
	return RedisEngine, nil
}

func NewRedisEngine(cfg *config.RedisCfg) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		Username: cfg.Username,
		DB:       cfg.DB,
	})
	ctx := redisClient.Context()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func NewDbEngine(cfg *config.MysqlCfg) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DBSetup() error {
	_, err := DefaultDB()
	if err != nil {
		return err
	}
	_, err = DefaultRedis()
	if err != nil {
		return err
	}
	return nil
}
