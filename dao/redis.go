package dao

import (
	"gin-api/exceptions"
	"gin-api/system"
	"time"

	"github.com/go-redis/redis"
)

func RedisConnect(config *system.RedisConfig) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.Host,
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	_, err := RedisClient.Ping().Result()

	exceptions.Error422(err)

	return
}
