package routes

import (
	"project-root/middlewares"
	"project-root/modules/auth/providers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RegisterRoutes(rg *gin.RouterGroup, authProvider *providers.Provider, rl *redis_rate.Limiter) {
	authRoutes := rg.Group("/auth")

	authRoutes.Use(middlewares.RateLimiter(rl, redis_rate.PerMinute(10)))

	authRoutes.POST("/login", authProvider.AuthController.Login)
	authRoutes.POST("/register", authProvider.AuthController.Register)
	authRoutes.POST("/refresh", authProvider.AuthController.Refresh)
	authRoutes.POST("/logout", authProvider.AuthController.Logout)
}
