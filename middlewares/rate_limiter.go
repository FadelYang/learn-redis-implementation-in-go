package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RateLimiter(limiter *redis_rate.Limiter, limitNumber redis_rate.Limit) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		ip := c.ClientIP()
		key := "rl:ip:" + ip

		res, err := limiter.Allow(ctx, key, limitNumber)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "rate limiter error"})
			return
		}

		if res.Allowed == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"errors": "too many request"})
			return
		}

		c.Next()
	}
}
