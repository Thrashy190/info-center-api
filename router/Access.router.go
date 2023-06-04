package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccessRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {

	accessService := service.AccessServiceImpl{DB: db}
	accessController := controllers.AccessController{AccessService: accessService}

	publicRouterAccess := publicRouter.Group("/access")
	{
		publicRouterAccess.POST("/students", accessController.CreateAccessStudents)
		publicRouterAccess.POST("/employees", accessController.CreateAccessEmployees)
	}

}
