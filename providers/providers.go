package providers

import (
	bookProvider "project-root/modules/books/providers"
	exProvider "project-root/modules/examples/providers"
	userProvider "project-root/modules/users/providers"

	"gorm.io/gorm"
)

type Providers struct {
	Examples *exProvider.Provider
	Users    *userProvider.Provider
	Books    *bookProvider.Provider
}

func Init(db *gorm.DB) *Providers {
	return &Providers{
		Examples: exProvider.NewProvider(db),
		Users:    userProvider.NewProvider(db),
		Books:    bookProvider.NewProvider(db),
	}
}
