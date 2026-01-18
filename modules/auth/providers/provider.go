package providers

import (
	"project-root/modules/auth/controller"
	authRepo "project-root/modules/auth/repository"
	userRepo "project-root/modules/users/repository"

	"project-root/modules/auth/service"
	userService "project-root/modules/users/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Provider struct {
	AuthController *controller.AuthController
}

func NewProvider(db *gorm.DB, redisClient *redis.Client) *Provider {
	repo := authRepo.NewAuthRepository(db)
	userRepo := userRepo.NewuserRepository(db)

	service := service.NewAuthService(repo, userRepo, redisClient)
	userService := userService.NewUserService(userRepo)

	controller := controller.NewAuthController(service, userService)

	return &Provider{
		AuthController: controller,
	}
}
