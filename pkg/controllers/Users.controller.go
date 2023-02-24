package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService service.UsersServiceImpl
}

func (c *UsersController) GetUsersByPage(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	offset := (page - 1) * 10
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Page"})
		return
	}
	users, err := c.UsersService.GetUsersByPage(int(offset))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": users})
}

func (c *UsersController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := c.UsersService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": user})
}

func (c *UsersController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := c.UsersService.UpdateUser(uint(id), user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": user})
}

func (c *UsersController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.UsersService.DeleteUser(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "User deleted"})
}
