package providers

import (
	"project-root/modules/books/controller"
	"project-root/modules/books/repository"
	"project-root/modules/books/services"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Provider struct {
	BookController *controller.BookController
}

func NewProvider(db *gorm.DB, redisClient *redis.Client) *Provider {
	repo := repository.NewBookRepository(db)
	service := services.NewBookService(repo, redisClient)
	controller := controller.NewBookController(service)

	return &Provider{
		BookController: controller,
	}
}
