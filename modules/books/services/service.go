package services

import (
	"project-root/modules/books/dto"
	"project-root/modules/books/repository"
)

type BookService interface {
	GetBooks() ([]dto.BookDTO, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (s *bookService) GetBooks() ([]dto.BookDTO, error) {
	books, err := s.bookRepository.GetAll()
	if err != nil {
		return nil, err
	}

	result := make([]dto.BookDTO, 0, len(books))
	for _, i := range books {
		result = append(result, dto.BookDTO{
			ID:          i.ID,
			Title:       i.Title,
			Description: i.Description,
			Author:      i.Author,
			Publisher:   i.Publisher,
		})
	}

	return result, nil
}
