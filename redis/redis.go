package redis

import (
	"gin/global"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	rds := redis.NewClient(&redis.Options{
		Addr:     global.REDIS_ADDR,
		Password: "",
		DB:       0,
	})

	global.RedisDB = rds
}
