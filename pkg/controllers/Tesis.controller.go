package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type TesisController struct {
	TesisService service.TesisServiceImpl
}

func (c *TesisController) CreateTesis(ctx *gin.Context) {
	var tesis models.Tesis
	if err := ctx.ShouldBindJSON(&tesis); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.TesisService.CreateTesis(&tesis); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": tesis})
}

func (c *TesisController) GetTesisByPage(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	offset := (page - 1) * 10
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Page"})
		return
	}
	tesiss, err := c.TesisService.GetTesisByPage(int(offset))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": tesiss})
}

func (c *TesisController) GetTesisByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	tesis, err := c.TesisService.GetTesissByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": tesis})
}

func (c *TesisController) UpdateTesis(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var tesis models.Tesis
	if err := ctx.ShouldBindJSON(&tesis); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := c.TesisService.UpdateTesis(uint(id), tesis); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": tesis})
}

func (c *TesisController) DeleteTesis(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.TesisService.DeleteTesis(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "Tesis deleted"})
}
