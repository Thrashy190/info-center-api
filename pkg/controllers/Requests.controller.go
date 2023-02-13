package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type RequestsController struct {
	RequestService service.RequestServiceImpl
}

func (c *RequestsController) CreateRequest(ctx *gin.Context) {
	var request models.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.RequestService.CreateRequest(&request); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": request})
}

func (c *RequestsController) GetRequestsByPage(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	offset := (page - 1) * 10
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Page"})
		return
	}
	requests, err := c.RequestService.GetRequestsByPage(int(offset))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": requests})
}

func (c *RequestsController) GetRequestByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	request, err := c.RequestService.GetRequestByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": request})
}

func (c *RequestsController) UpdateRequest(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.RequestService.UpdateRequest(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "Request updated successfully"})
}
