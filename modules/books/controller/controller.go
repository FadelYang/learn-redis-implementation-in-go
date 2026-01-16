package controller

import (
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/books/dto"
	"project-root/modules/books/services"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

// @Tags books
// @Summary get books
// @Description get all books
// @Accept json
// @Produce json
// @Success 200 {object} common.BaseResponse[dto.BookDTO]
// @Router /books [get]
func (c *BookController) GetAll(ctx *gin.Context) {
	books, err := c.bookService.GetBooks()
	if err != nil {
		log.Fatalf("failed to get books: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "failed to get books data"})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse[[]dto.BookDTO]{
		Status:  http.StatusOK,
		Message: "success get books data",
		Data:    books,
	})
}
