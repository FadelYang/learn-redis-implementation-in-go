package service

import (
	"context"
	"errors"
	"project-root/modules/auth/dto"
	authRepository "project-root/modules/auth/repository"
	userRepository "project-root/modules/users/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var (
	accessSecret  = []byte("ACCESS_SECRET_KEY")
	refreshSecret = []byte("REFRESH_SECRET_KEY")
)

type AuthService interface {
	Login(ctx context.Context, loginForm dto.LoginDTO) (dto.LoginResponse, error)
	RefreshLogin(ctx context.Context, refreshToken string) (string, error)
}

type authService struct {
	authRepository authRepository.AuthRepository
	userRepository userRepository.UserRepository
	redisClient    *redis.Client
}

func NewAuthService(
	authRepository authRepository.AuthRepository,
	userRepository userRepository.UserRepository,
	redisClient *redis.Client,
) AuthService {
	return &authService{
		authRepository: authRepository,
		userRepository: userRepository,
		redisClient:    redisClient,
	}
}

func (s *authService) Login(ctx context.Context, loginForm dto.LoginDTO) (dto.LoginResponse, error) {
	user, err := s.userRepository.FindByEmail(loginForm.Email)
	if err != nil {
		return dto.LoginResponse{}, authRepository.ErrLogin
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(loginForm.Password),
	); err != nil {
		return dto.LoginResponse{}, authRepository.ErrLogin
	}

	accessToken, err := s.generateAccessToken(user.ID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	refreshToken, err := s.generateRefreshToken(ctx, user.ID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) generateAccessToken(userID uuid.UUID) (string, error) {
	claims := dto.AccessClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(accessSecret)
}

func (s *authService) generateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	jti := uuid.NewString()

	claims := dto.RefreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	err = s.redisClient.Set(
		ctx,
		"refresh:"+jti,
		userID.String(),
		7*24*time.Hour,
	).Err()
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (s *authService) RefreshLogin(ctx context.Context, refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&dto.RefreshClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return refreshSecret, nil
		},
	)
	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims := token.Claims.(*dto.RefreshClaims)

	return s.generateAccessToken(claims.UserID)
}
