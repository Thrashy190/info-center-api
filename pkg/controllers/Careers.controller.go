package controllers

import (
	"strconv"

	"github.com/Thrashy190/info-center-api/pkg/models"
	service "github.com/Thrashy190/info-center-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type CareersController struct {
	CareerService service.CareersServiceImpl
}

//	@Summary		Create a career
//	@Description	Create a career
//	@Tags			Careers
//	@Produce		json
//	@Param			career	body		models.CareersForSwagger	true	"Career JSON"
//	@Success		200		{object}	models.Careers
//	@Failure		400		"Invalid request payload"
//	@Failure		500		"Internal Server Error"
//	@Router			/protected/careers/career [post]
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

//	@Summary		Get all Careers
//	@Description	Get all the list of careers
//	@Tags			Careers
//	@Produce		json
//	@Success		200	{object}	[]models.Careers
//	@Failure		400	"Invalid request payload"
//	@Failure		500	"Internal Server Error"
//	@Router			/public/careers/ [get]
func (c *CareersController) GetCareers(ctx *gin.Context) {

	careers, err := c.CareerService.GetCareers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": careers})
}

//	@Summary		Update Career
//	@Description	Update a career
//	@Tags			Careers
//	@Produce		json
//	@Param			id	path		string	true	"ID"
//	@Success		200	{object}	models.Careers
//	@Failure		400	"Invalid request payload"
//	@Failure		500	"Internal Server Error"
//	@Router			/protected/careers/{id} [put]
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

//	@Summary		Delete Career
//	@Description	Delete career
//	@Tags			Careers
//	@Produce		json
//	@Param			id	path	string	true	"ID"
//	@Success		200	"Career deleted"
//	@Failure		400	"Invalid request payload"
//	@Failure		500	"Internal Server Error"
//	@Router			/protected/careers/{id} [delete]
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
