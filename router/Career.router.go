package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CareersRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {
	careerService := service.CareersServiceImpl{DB: db}
	careerController := controllers.CareersController{CareerService: careerService}

	publicRouterCarrers := publicRouter.Group("/careers")
	{
		publicRouterCarrers.GET("/", careerController.GetCareers)
	}

	protectedRouterCarrers := protectedRouter.Group("/careers")
	{
		protectedRouterCarrers.POST("/career", careerController.CreateCareer)
		protectedRouterCarrers.PUT("/:id", careerController.UpdateCareer)
		protectedRouterCarrers.DELETE("/:id", careerController.DeleteCareer)
	}

}
