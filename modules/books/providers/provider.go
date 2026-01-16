package providers

import (
	"project-root/modules/books/controller"
	"project-root/modules/books/repository"
	"project-root/modules/books/services"

	"gorm.io/gorm"
)

type Provider struct {
	BookController *controller.BookController
}

func NewProvider(db *gorm.DB) *Provider {
	repo := repository.NewBookRepository(db)
	service := services.NewBookService(repo)
	controller := controller.NewBookController(service)

	return &Provider{
		BookController: controller,
	}
}
