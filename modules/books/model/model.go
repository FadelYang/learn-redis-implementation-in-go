package model

import (
	"time"

	"github.com/google/uuid"
)

type BookModel struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `gorm:"type:varchar(50);unique;not null"`
	Description string    `gorm:"type:text;not null"`
	Author      string    `gorm:"type:varchar(75);not null"`
	Publisher   string    `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (BookModel) TableName() string {
	return "books"
}
