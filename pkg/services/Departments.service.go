package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type DepartmentService interface {
	GetDepartments() ([]models.Departments, error)
	CreateDepartment(*models.Departments) error
	DeleteDepartment(id uint) error
	UpdateDepartment(id uint, department models.Departments) error
}

type DepartmentServiceImpl struct {
	DB *gorm.DB
}

func (s *DepartmentServiceImpl) GetDepartments() ([]models.Departments, error) {
	var department []models.Departments
	if err := s.DB.Find(&department).Error; err != nil {
		return nil, err
	}
	return department, nil
}

func (s *DepartmentServiceImpl) CreateDepartment(department *models.Departments) error {
	if err := s.DB.Create(department).Error; err != nil {
		return err
	}
	return nil
}

func (s *DepartmentServiceImpl) UpdateDepartment(id uint, department models.Departments) error {
	if err := s.DB.Model(&department).Where("id = ?", id).Updates(&department).Error; err != nil {
		return err
	}
	return nil
}

func (s *DepartmentServiceImpl) DeleteDepartment(id uint) error {
	var department models.Departments
	if err := s.DB.Where("id = ?", id).Delete(&department).Error; err != nil {
		return err
	}
	return nil
}
