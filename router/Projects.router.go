package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProjectsRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {

	projectService := service.ProjectsServiceImpl{DB: db}
	projectController := controllers.ProjectsController{ProjectService: projectService}

	publicRouterProjects := publicRouter.Group("/projects")
	{
		publicRouterProjects.GET("/project/:id", projectController.GetProjectByID)
		publicRouterProjects.GET("/:page", projectController.GetProjectsByPage)
	}
	protectedRouterProjects := protectedRouter.Group("/projects")
	{
		protectedRouterProjects.POST("/project", projectController.CreateProject)
		protectedRouterProjects.PUT("project/:id", projectController.UpdateProject)
		protectedRouterProjects.DELETE("/project/:id", projectController.DeleteProject)
	}

}
