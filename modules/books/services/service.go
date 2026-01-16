package services

import (
	"context"
	"encoding/json"
	"log"
	"project-root/modules/books/dto"
	"project-root/modules/books/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type BookService interface {
	GetBooks() ([]dto.BookDTO, error)
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

func (s *bookService) GetBooks() ([]dto.BookDTO, error) {
	ctx := context.Background()
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

	bytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	if err := s.redisClient.Set(ctx, cacheKey, bytes, time.Minute).Err(); err != nil {
		return nil, err
	} else {
		log.Println("sucess store data to redis")
	}

	return result, nil
}
