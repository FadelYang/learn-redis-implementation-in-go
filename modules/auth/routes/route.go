package routes

import (
	"project-root/modules/auth/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, authProvider *providers.Provider) {
	authRoutes := rg.Group("/auth")

	authRoutes.POST("/login", authProvider.AuthController.Login)
	authRoutes.POST("/register", authProvider.AuthController.Register)
}
