package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(gin *gin.Engine, db *gorm.DB) {

	publicRouter := gin.Group("/public")
	protectedRouter := gin.Group("/protected")

	BooksRouter(publicRouter, protectedRouter, db)
	RequestsRouter(publicRouter, protectedRouter, db)
	CareersRouter(publicRouter, protectedRouter, db)

	ProjectsRouter(publicRouter, protectedRouter, db)

}
