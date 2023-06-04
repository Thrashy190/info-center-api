package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	DepartmentService service.DepartmentServiceImpl
}

func (c *DepartmentController) CreateDepartment(ctx *gin.Context) {
	var department models.Departments
	if err := ctx.ShouldBindJSON(&department); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.DepartmentService.CreateDepartment(&department); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": department})
}

func (c *DepartmentController) GetDepartments(ctx *gin.Context) {

	department, err := c.DepartmentService.GetDepartments()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": department})
}

func (c *DepartmentController) UpdateDepartment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var department models.Departments
	if err := ctx.ShouldBindJSON(&department); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := c.DepartmentService.UpdateDepartment(uint(id), department); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": department})
}

func (c *DepartmentController) DeleteDepartment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.DepartmentService.DeleteDepartment(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "Department deleted"})
}
