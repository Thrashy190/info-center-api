package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type AccessService interface {
	CreateAccessStudents(*models.AccessStudents) error
	CreateAccessEmployees(*models.AccessEmployees) error
}

type AccessServiceImpl struct {
	DB *gorm.DB
}

func (s *AccessServiceImpl) CreateAccessStudents(access *models.AccessStudents) error {
	if err := s.DB.Create(access).Error; err != nil {
		return err
	}
	return nil
}

func (s *AccessServiceImpl) CreateAccessEmployees(access *models.AccessEmployees) error {
	if err := s.DB.Create(access).Error; err != nil {
		return err
	}
	return nil
}
