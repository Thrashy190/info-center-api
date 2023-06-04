package router

import (
	"github.com/Thrashy190/info-center-api/pkg/controllers"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DepartmentRouter(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, db *gorm.DB) {
	departmentService := service.DepartmentServiceImpl{DB: db}
	departmentController := controllers.DepartmentController{DepartmentService: departmentService}

	publicRouterDepartments := publicRouter.Group("/departments")
	{
		publicRouterDepartments.GET("/", departmentController.GetDepartments)
	}

	protectedRouterDepartments := protectedRouter.Group("/departments")
	{
		protectedRouterDepartments.POST("/department", departmentController.CreateDepartment)
		protectedRouterDepartments.PUT("/:id", departmentController.UpdateDepartment)
		protectedRouterDepartments.DELETE("/:id", departmentController.DeleteDepartment)
	}

}
