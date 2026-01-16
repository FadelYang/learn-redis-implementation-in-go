package config

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func InitRedis() *redis.Client {
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		panic("failed to connect redis: " + err.Error())
	}

	return rdb
}
