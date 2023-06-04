package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type CareersService interface {
	GetCareers() ([]models.Careers, error)
	CreateCareer(*models.Careers) error
	DeleteCareer(id uint) error
	UpdateCareer(id uint, Career models.Careers) error
}

type CareersServiceImpl struct {
	DB *gorm.DB
}

func (s *CareersServiceImpl) GetCareers() ([]models.Careers, error) {
	var careers []models.Careers
	if err := s.DB.Find(&careers).Error; err != nil {
		return nil, err
	}
	return careers, nil
}

func (s *CareersServiceImpl) CreateCareer(career *models.Careers) error {
	if err := s.DB.Create(career).Error; err != nil {
		return err
	}
	return nil
}

func (s *CareersServiceImpl) UpdateCareer(id uint, career models.Careers) error {
	if err := s.DB.Model(&career).Where("id = ?", id).Updates(&career).Error; err != nil {
		return err
	}
	return nil
}

func (s *CareersServiceImpl) DeleteCareer(id uint) error {
	var career models.Careers
	if err := s.DB.Where("id = ?", id).Delete(&career).Error; err != nil {
		return err
	}
	return nil
}
