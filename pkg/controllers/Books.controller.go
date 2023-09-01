package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type BooksController struct {
	BookService service.BooksServiceImpl
}

// @Summary		Create a book
// @Description	Create a book
// @Tags			Books
// @Produce		json
// @Param			book	body		models.BookForSwagger	true	"Book JSON"
// @Success		200		{object}	models.Book
// @Failure		400		"Invalid request payload"
// @Failure		500		"Internal Server Error"
// @Router			/protected/books/book [post]
func (c *BooksController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.BookService.CreateBook(&book); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"book": book})
}

// @Summary		Get Books by Page
// @Description	Pagination for books
// @Tags			Books
// @Produce		json
// @Param			page	path		string	true	"Page ID"
// @Success		200		{object}	[]models.Book
// @Failure		400		"Invalid request payload"
// @Failure		500		"Internal Server Error"
// @Router			/public/books/{page} [get]
func (c *BooksController) GetBooksByPage(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Page"})
		return
	}

	paginationData, err := c.BookService.GetBooksByPage(int(page))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": paginationData})
}

// @Summary		Get Books by ID
// @Description	Get only one book by id
// @Tags			Books
// @Produce		json
// @Param			id	path		string	true	"ID"
// @Success		200	{object}	models.Book
// @Failure		400	"Invalid request payload"
// @Failure		500	"Internal Server Error"
// @Router			/public/books/book/{id} [get]
func (c *BooksController) GetBooksByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	book, err := c.BookService.GetBooksByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": book})
}

// @Summary		Update book
// @Description	Update book by id
// @Tags			Books
// @Produce		json
// @Param			id	path		string	true	"ID"
// @Success		200	{object}	models.Book
// @Failure		400	"Invalid request payload"
// @Failure		500	"Internal Server Error"
// @Router			/protected/books/book/{id} [put]
func (c *BooksController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := c.BookService.UpdateBook(uint(id), book); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": book})
}

// @Summary		Delete book
// @Description	Delete book by id
// @Tags			Books
// @Produce		json
// @Param			id	path	string	true	"ID"
// @Success		200	"User deleted"
// @Failure		400	"Invalid request payload"
// @Failure		500	"Internal Server Error"
// @Router			/protected/books/book/{id} [delete]
func (c *BooksController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.BookService.DeleteBook(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "User deleted"})
}

// 128+64+32+16+8+4+2+1
