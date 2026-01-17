package model

import "github.com/google/uuid"

type Example struct {
	UUID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"not null;unique"`
}
