package routes

import (
	"project-root/modules/books/providers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, bookProvider *providers.Provider) {
	bookRoutes := rg.Group("/books")

	bookRoutes.GET("", bookProvider.BookController.GetAll)
}
