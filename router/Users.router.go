package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRouter(protectedRouter *gin.RouterGroup, db *gorm.DB) {

	userService := service.UsersServiceImpl{DB: db}
	userController := controllers.UsersController{UsersService: userService}

	protectedRouterUser := protectedRouter.Group("/users")
	{
		protectedRouterUser.GET("/user/:id", userController.GetUserByID)
		protectedRouterUser.GET("/:page", userController.GetUsersByPage)
		protectedRouterUser.DELETE("/user/:id", userController.DeleteUser)
		protectedRouterUser.PUT("/user/:id", userController.UpdateUser)
	}

}
