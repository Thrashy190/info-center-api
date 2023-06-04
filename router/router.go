package router

import (
	"github.com/Thrashy190/info-center-api/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gin *gin.Engine, db *gorm.DB) {

	requireAuthMiddleware := middlewares.RequireAuthImpl{DB: db}

	publicRouter := gin.Group("/public")
	protectedRouter := gin.Group("/protected")

	protectedRouter.Use(requireAuthMiddleware.RequireAuthMiddleware)

	AccessRouter(publicRouter, protectedRouter, db)
	DepartmentRouter(publicRouter, protectedRouter, db)
	BooksRouter(publicRouter, protectedRouter, db)
	RequestsRouter(publicRouter, protectedRouter, db)
	CareersRouter(publicRouter, protectedRouter, db)
	TesisRouter(publicRouter, protectedRouter, db)
	UsersRouter(protectedRouter, db)
	ProjectsRouter(publicRouter, protectedRouter, db)
	AuthRouter(publicRouter, protectedRouter, db)

}
