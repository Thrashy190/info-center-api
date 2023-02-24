package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TesisRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {

	tesisService := service.TesisServiceImpl{DB: db}
	tesisController := controllers.TesisController{TesisService: tesisService}

	publicRouterTesis := publicRouter.Group("/tesis")
	{
		publicRouterTesis.GET("/:page", tesisController.GetTesisByPage)
		publicRouterTesis.GET("/tesis/:id", tesisController.GetTesisByID)
	}

	protectedRouterTesis := protectedRouter.Group("/tesis")
	{
		protectedRouterTesis.POST("/tesis", tesisController.CreateTesis)
		protectedRouterTesis.DELETE("/tesis/:id", tesisController.DeleteTesis)
		protectedRouterTesis.PUT("/tesis/:id", tesisController.UpdateTesis)
	}

}
