package providers

import (
	authProvider "project-root/modules/auth/providers"
	bookProvider "project-root/modules/books/providers"
	exProvider "project-root/modules/examples/providers"
	userProvider "project-root/modules/users/providers"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Providers struct {
	Examples *exProvider.Provider
	Users    *userProvider.Provider
	Books    *bookProvider.Provider
	Auth     *authProvider.Provider
}

func Init(db *gorm.DB, redisClient *redis.Client) *Providers {
	return &Providers{
		Examples: exProvider.NewProvider(db),
		Users:    userProvider.NewProvider(db),
		Books:    bookProvider.NewProvider(db, redisClient),
		Auth:     authProvider.NewProvider(db, redisClient),
	}
}
