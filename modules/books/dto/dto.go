package dto

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id" example:"a53515e3-5a7f-440b-82f6-3d84ac7ce746"`
	Title       string    `json:"title" example:"Sang Mentari Di Malam Hari"`
	Description string    `json:"description" example:"di sebuah malam yang gelap, di mana kah sang mentari terlelap?"`
	Author      string    `json:"author" example:"Michael Van Brave"`
	Publisher   string    `json:"publisher" example:"Penerbit Kakamuda"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
