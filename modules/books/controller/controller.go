package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/books/dto"
	"project-root/modules/books/repository"
	"project-root/modules/books/services"
	"project-root/modules/users/service"
	"project-root/tools"

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
	books, err := c.bookService.GetBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "failed to get books data"})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse[[]dto.BookDTO]{
		Status:  http.StatusOK,
		Message: "success get books data",
		Data:    books,
	})
}

// @Tags books
// @Summary get books
// @Description get all books
// @Accept json
// @Produce json
// @Success 201 {object} common.BaseResponse[dto.BookDTO]
// @Router /books [post]
// @Param request body dto.BookDTO true "request bodu for create a book [RAW]"
func (c *BookController) Create(ctx *gin.Context) {
	var book dto.BookDTO
	if err := ctx.ShouldBindBodyWithJSON(&book); err != nil {
		log.Printf("failed to create a book: %v", err)

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": fmt.Sprintf("%s: %s", repository.ErrCreateBookValidate,
					err.Error()),
			},
		)
		return
	}

	createdBook, err := c.bookService.Create(ctx, book)
	if err != nil {
		log.Printf("Failed to create user: %v", err)

		var vErr *tools.ValidationError
		if errors.As(err, &vErr) {
			ctx.JSON(http.StatusConflict, vErr)
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", service.ErrCreateUserValidate, err.Error())})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		common.BaseResponse[dto.BookDTO]{
			Status:  http.StatusCreated,
			Message: "book created successfully",
			Data:    createdBook,
		},
	)
}
