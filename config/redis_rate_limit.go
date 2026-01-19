package config

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

func InitRedisRateLimit(redisClient *redis.Client) *redis_rate.Limiter {
	return redis_rate.NewLimiter(redisClient)
}
