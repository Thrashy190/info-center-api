package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type TesissService interface {
	GetTesisByPage(page uint) ([]models.Tesis, error)
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

func (s *TesisServiceImpl) GetTesisByPage(page int) ([]models.Tesis, error) {
	var tesis []models.Tesis
	if err := s.DB.Limit(10).Offset(page).Preload("Careers").Find(&tesis).Error; err != nil {
		return nil, err
	}
	return tesis, nil
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
