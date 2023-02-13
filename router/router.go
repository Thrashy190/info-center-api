package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {

	bookService := service.BooksServiceImpl{DB: db}
	bookController := controllers.BooksController{BookService: bookService}

	requestService := service.RequestServiceImpl{DB: db}
	requestController := controllers.RequestsController{RequestService: requestService}

	careerService := service.CareersServiceImpl{DB: db}
	careerController := controllers.CareersController{CareerService: careerService}

	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.POST("/book", bookController.CreateBook)
			books.GET("/:page", bookController.GetBooksByPage)
			books.DELETE("/book/:id", bookController.DeleteBook)
			books.GET("/book/:id", bookController.GetBooksByID)
			books.PUT("/book/:id", bookController.UpdateBook)
		}
		requests := v1.Group("/requests")
		{
			requests.POST("/request", requestController.CreateRequest)
			requests.GET("/:page", requestController.GetRequestsByPage)
			requests.GET("/request/:id", requestController.GetRequestByID)
			requests.PUT("/request/:id", requestController.UpdateRequest)
		}
		carrers := v1.Group("/careers")
		{
			carrers.POST("/career", careerController.CreateCareer)
			carrers.GET("/", careerController.GetCareers)
			carrers.PUT("/:id", careerController.UpdateCareer)
			carrers.DELETE("/:id", careerController.DeleteCareer)
		}
	}
}
