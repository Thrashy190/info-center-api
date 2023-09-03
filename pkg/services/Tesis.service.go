package service

import (
	"math"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type TesissService interface {
	GetTesisByPage(page uint) (*models.PaginationDataTesis, error)
	GetTesisByID(id uint) (*models.Tesis, error)
	CreateTesis(*models.Tesis) error
	DeleteTesis(id uint) error
	UpdateTesis(id uint, tesis models.Tesis) error
}

type TesisServiceImpl struct {
	DB *gorm.DB
}

func (s *TesisServiceImpl) CreateTesis(tesis *models.Tesis) error {
	if err := s.DB.Create(tesis).Error; err != nil {
		return err
	}
	return nil
}

func (s *TesisServiceImpl) GetTesisByPage(page int) (*models.PaginationDataTesis, error) {

	perPage := 4

	//Count total tesis
	var totalRows int64
	if err := s.DB.Model(&models.Tesis{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}
	totalPages := math.Ceil(float64(totalRows) / float64(perPage))

	offset := (page - 1) * perPage

	//Get pagination data
	var tesis []models.Tesis
	if err := s.DB.Limit(perPage).Offset(offset).Preload("Careers").Find(&tesis).Error; err != nil {
		return nil, err
	}

	var paginationData = models.PaginationDataTesis{
		Tesis:        &tesis,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   int(totalPages),
	}

	return &paginationData, nil
}

func (s *TesisServiceImpl) GetTesissByID(id uint) (*models.Tesis, error) {
	var tesis models.Tesis
	if err := s.DB.Where("id = ?", id).Preload("Careers").First(&tesis).Error; err != nil {
		return nil, err
	}
	return &tesis, nil
}

func (s *TesisServiceImpl) UpdateTesis(id uint, tesis models.Tesis) error {
	if err := s.DB.Model(&tesis).Where("id = ?", id).Updates(&tesis).Error; err != nil {
		return err
	}
	return nil
}

func (s *TesisServiceImpl) DeleteTesis(id uint) error {
	var tesis models.Tesis
	if err := s.DB.Where("id = ?", id).Delete(&tesis).Error; err != nil {
		return err
	}
	return nil
}
