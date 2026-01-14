package seeders

import (
	"fmt"

	"project-root/modules/books/dto"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookSeeder struct{}

func (s BookSeeder) Run(db *gorm.DB) error {
	gofakeit.Seed(0)

	for range 1000 {
		newUUID, _ := uuid.Parse(gofakeit.UUID())
		book := dto.Book{
			ID:          newUUID,
			Title:       gofakeit.BookTitle(),
			Description: gofakeit.ProductDescription(),
			Author:      gofakeit.BookAuthor(),
			Publisher:   fmt.Sprintf("Penerbit %s", gofakeit.AppName()),
		}

		if err := db.
			Clauses(clause.OnConflict{
				DoNothing: true,
			}).
			Create(&book).Error; err != nil {
			return err
		}
	}

	return nil
}
