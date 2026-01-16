package repository

import (
	"project-root/modules/books/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]model.BookModel, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) GetAll() ([]model.BookModel, error) {
	var books []model.BookModel

	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
