package router

import (
	"github.com/Thrashy190/info-center-api/pkg/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gin *gin.Engine, db *gorm.DB) {

	requireAuthMiddleware := middlewares.RequireAuthImpl{DB: db}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}                    // Replace with your frontend URL
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUTS", "DELETE"} // Add the HTTP methods your app uses
	gin.Use(cors.New(config))

	gin.Use(middlewares.CORSMiddleware())

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
