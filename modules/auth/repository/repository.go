package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrLogin    = errors.New("login failed: incorrect email or password")
	ErrRegister = errors.New("register failed")
)

type AuthRepository interface{}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
