package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RateLimiter(limiter *redis_rate.Limiter, limitNumber redis_rate.Limit) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		ip := c.ClientIP()
		if ip == "::1" || ip == "0:0:0:0:0:0:0:1" {
			ip = "127.0.0.1"
		}

		route := c.FullPath()

		ip = strings.ReplaceAll(ip, ":", "_")
		route = strings.ReplaceAll(route, "/", "_")

		key := fmt.Sprintf("%s:%s", route, ip)

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
