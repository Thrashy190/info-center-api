package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type BooksController struct {
	BookService service.BooksServiceImpl
}

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

	ctx.JSON(200, gin.H{"data": book})
}

func (c *BooksController) GetBooksByPage(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	offset := (page - 1) * 10
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Page"})
		return
	}
	books, err := c.BookService.GetBooksByPage(int(offset))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": books})
}

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
