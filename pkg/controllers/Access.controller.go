package controllers

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type AccessController struct {
	AccessService service.AccessServiceImpl
}

func (c *AccessController) CreateAccessStudents(ctx *gin.Context) {
	var access models.AccessStudents

	if ctx.Bind(&access) != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.AccessService.CreateAccessStudents(&access); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to create access", "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Access created successfully"})
}

func (c *AccessController) CreateAccessEmployees(ctx *gin.Context) {
	var access models.AccessEmployees

	if ctx.Bind(&access) != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.AccessService.CreateAccessEmployees(&access); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to create access", "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Access created successfully"})
}
