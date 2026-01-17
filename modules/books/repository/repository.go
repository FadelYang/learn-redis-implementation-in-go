package repository

import (
	"errors"
	"project-root/modules/books/model"
	"project-root/tools"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var (
	ErrBookNotFound       = errors.New("book not found")
	ErrCreateBookValidate = errors.New("failed to create a book")
	ErrUpdateBookVaildate = errors.New("failed to update a book")
	ErrDeleteABook        = errors.New("failed to delete a book")
	ErrDuplicateBookTitle = errors.New("book with that title are already exists")
)

type BookRepository interface {
	GetAll() ([]model.BookModel, error)
	Create(book model.BookModel) (model.BookModel, error)
	Update(book model.BookModel) (model.BookModel, error)
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (model.BookModel, error)
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

func (r *bookRepository) Create(book model.BookModel) (model.BookModel, error) {
	ve := tools.NewValidationError()

	if err := r.db.Create(&book).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr != nil {
			switch pgErr.ConstraintName {
			case "books_title_key":
				ve.Add("title", ErrDuplicateBookTitle.Error())

				return model.BookModel{}, ve
			default:
				return model.BookModel{}, err
			}
		}
		return model.BookModel{}, err
	}

	return book, nil
}

func (r *bookRepository) Update(book model.BookModel) (model.BookModel, error) {
	ve := tools.NewValidationError()

	if err := r.db.Model(&book).Updates(map[string]any{
		"title":       book.Title,
		"description": book.Description,
		"author":      book.Author,
	}).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr != nil {
			switch pgErr.ConstraintName {
			case "books_title_key":
				ve.Add("title", ErrDuplicateBookTitle.Error())
				return model.BookModel{}, ve
			default:
				return model.BookModel{}, err
			}
		}
		return model.BookModel{}, err
	}

	return book, nil
}

func (r *bookRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&model.BookModel{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) GetByID(id uuid.UUID) (model.BookModel, error) {
	var book model.BookModel
	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		return model.BookModel{}, err
	}

	return book, nil
}
