package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type CareersController struct {
	CareerService service.CareersServiceImpl
}

func (c *CareersController) CreateCareer(ctx *gin.Context) {
	var career models.Careers
	if err := ctx.ShouldBindJSON(&career); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.CareerService.CreateCareer(&career); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": career})
}

func (c *CareersController) GetCareers(ctx *gin.Context) {

	careers, err := c.CareerService.GetCareers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": careers})
}

func (c *CareersController) UpdateCareer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var career models.Careers
	if err := ctx.ShouldBindJSON(&career); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := c.CareerService.UpdateCareer(uint(id), career); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": career})
}

func (c *CareersController) DeleteCareer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.CareerService.DeleteCareer(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "Career deleted"})
}
