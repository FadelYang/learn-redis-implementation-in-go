package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type LoginDTO struct {
	Email    string `json:"email" example:"fadelanumah@gmail.com"`
	Password string `json:"password" example:"Secretpassword@123"`
}

type RegisterDTO struct {
	Email    string `json:"email" example:"fadelanumah@gmail.com"`
	Password string `json:"password" example:"Secretpassword@123"`
	Username string `json:"username" example:"fadelanumah"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"Grgws!23412sA"`
	RefreshToken string `json:"refresh_token" example:"Grgws!23412sA"`
}

type AccessClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}
