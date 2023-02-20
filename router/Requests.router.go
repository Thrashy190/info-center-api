package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequestsRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {
	requestService := service.RequestServiceImpl{DB: db}
	requestController := controllers.RequestsController{RequestService: requestService}

	publicRouterRequests := publicRouter.Group("/requests")
	{
		publicRouterRequests.POST("/request", requestController.CreateRequest)
	}

	protectedRouterRequests := protectedRouter.Group("/requests")
	{
		protectedRouterRequests.GET("/:page", requestController.GetRequestsByPage)
		protectedRouterRequests.GET("/request/:id", requestController.GetRequestByID)
		protectedRouterRequests.PUT("/request/:id", requestController.UpdateRequest)
	}

}
