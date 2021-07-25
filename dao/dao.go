package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"gin-api/system"
)

var (
	RedisClient *redis.Client
	MysqlClient   *gorm.DB
	err         error
)

func Connect() {
	config := system.ProjectConfig

	MysqlConnect(config.Mysql)
	RedisConnect(config.Redis)
}

func Close() {
	RedisClient.Close()
}
