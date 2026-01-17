package services

import (
	"context"
	"encoding/json"
	"errors"
	"project-root/modules/books/dto"
	"project-root/modules/books/model"
	"project-root/modules/books/repository"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BookService interface {
	GetBooks(ctx context.Context) ([]dto.BookDTO, error)
	Create(ctx context.Context, form dto.BookDTO) (dto.BookDTO, error)
	Update(ctx context.Context, form dto.BookDTO, bookID uuid.UUID) (dto.BookDTO, error)
	Delete(ctx context.Context, bookID uuid.UUID) error
}

type bookService struct {
	bookRepository repository.BookRepository
	redisClient    *redis.Client
}

func NewBookService(
	bookRepository repository.BookRepository,
	redisClient *redis.Client,
) BookService {
	return &bookService{
		bookRepository: bookRepository,
		redisClient:    redisClient,
	}
}

func (s *bookService) GetBooks(ctx context.Context) ([]dto.BookDTO, error) {
	cacheKey := "books:all"

	val, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedBooks []dto.BookDTO
		if err := json.Unmarshal([]byte(val), &cachedBooks); err == nil {
			return cachedBooks, nil
		}
	}

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

	if bytes, err := json.Marshal(result); err == nil {
		_ = s.redisClient.Set(ctx, cacheKey, bytes, time.Minute).Err()
	}

	return result, nil
}

func (s *bookService) Create(ctx context.Context, book dto.BookDTO) (dto.BookDTO, error) {
	bookForm := model.BookModel{
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		Publisher:   book.Publisher,
	}

	createdBook, err := s.bookRepository.Create(bookForm)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateBookTitle) {
			return dto.BookDTO{}, repository.ErrDuplicateBookTitle
		}
		return dto.BookDTO{}, err
	}

	_ = s.redisClient.Del(ctx, "books:all").Err()

	return dto.BookDTO{
		ID:          createdBook.ID,
		Title:       createdBook.Title,
		Description: createdBook.Description,
		Publisher:   createdBook.Publisher,
		Author:      createdBook.Author,
	}, nil
}

func (s *bookService) Update(ctx context.Context, book dto.BookDTO, bookID uuid.UUID) (dto.BookDTO, error) {
	existingBook, err := s.bookRepository.GetByID(bookID)
	if err != nil {
		return dto.BookDTO{}, err
	}

	existingBook.Title = book.Title
	existingBook.Description = book.Description
	existingBook.Author = book.Author
	existingBook.Publisher = book.Publisher

	updatedBook, err := s.bookRepository.Update(existingBook)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateBookTitle) {
			return dto.BookDTO{}, repository.ErrDuplicateBookTitle
		}
		return dto.BookDTO{}, err
	}

	_ = s.redisClient.Del(ctx, "books:all").Err()

	return dto.BookDTO{
		ID:          updatedBook.ID,
		Title:       updatedBook.Title,
		Description: updatedBook.Description,
		Publisher:   updatedBook.Publisher,
		Author:      updatedBook.Author,
	}, nil
}

func (s *bookService) Delete(ctx context.Context, bookID uuid.UUID) error {
	existingBook, err := s.bookRepository.GetByID(bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.ErrBookNotFound
		}
		return err
	}

	_ = s.redisClient.Del(ctx, "books:all").Err()

	if err := s.bookRepository.Delete(existingBook.ID); err != nil {
		return err
	}

	return nil
}
