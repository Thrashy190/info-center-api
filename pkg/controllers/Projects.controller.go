package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type ProjectsController struct {
	ProjectService service.ProjectsServiceImpl
}

func (c *ProjectsController) CreateProject(ctx *gin.Context) {
	var project models.Projects
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.ProjectService.CreateProject(&project); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": project})
}

func (c *ProjectsController) GetProjectsByPage(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	offset := (page - 1) * 10
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Page"})
		return
	}
	projects, err := c.ProjectService.GetProjectsByPage(int(offset))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": projects})
}

func (c *ProjectsController) GetProjectByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	project, err := c.ProjectService.GetProjectsByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": project})
}

func (c *ProjectsController) UpdateProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var project models.Projects
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := c.ProjectService.UpdateProject(uint(id), project); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": project})
}

func (c *ProjectsController) DeleteProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.ProjectService.DeleteProject(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": "Project deleted"})
}
