package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BooksRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {

	bookService := service.BooksServiceImpl{DB: db}
	bookController := controllers.BooksController{BookService: bookService}

	publicRouterBooks := publicRouter.Group("/books")
	{
		publicRouterBooks.GET("/:page", bookController.GetBooksByPage)
		publicRouterBooks.GET("/book/:id", bookController.GetBooksByID)
	}

	protectedRouterBooks := protectedRouter.Group("/books")
	{
		protectedRouterBooks.POST("/book", bookController.CreateBook)
		protectedRouterBooks.DELETE("/book/:id", bookController.DeleteBook)
		protectedRouterBooks.PUT("/book/:id", bookController.UpdateBook)
	}

}
