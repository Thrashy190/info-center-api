package router

import (
	"github.com/Thrashy190/info-center-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {

	authService := auth.AuthenticationServiceImpl{DB: db}
	authController := auth.AuthenticationController{AuthService: authService}

	publicRouterAuth := publicRouter.Group("/auth")
	{
		publicRouterAuth.POST("/signup", authController.SignUp)
		publicRouterAuth.POST("/login", authController.Login)
	}

	protectedRouterAuth := protectedRouter.Group("/auth")
	{
		protectedRouterAuth.POST("/logout", authController.SignOut)
	}

}
