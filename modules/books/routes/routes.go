package routes

import (
	"project-root/middlewares"
	"project-root/modules/books/providers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RegisterRoutes(rg *gin.RouterGroup, bookProvider *providers.Provider, rl *redis_rate.Limiter) {
	bookRoutes := rg.Group("/books")

	bookRoutes.Use(middlewares.RateLimiter(rl, redis_rate.PerMinute(500)))

	bookRoutes.GET("", bookProvider.BookController.GetAll)
	bookRoutes.POST("", bookProvider.BookController.Create)
}
