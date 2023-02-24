package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type ProjectsService interface {
	GetProjectsByPage(page uint) ([]models.Projects, error)
	GetProjectByID(id uint) (*models.Projects, error)
	CreateProject(*models.Projects) error
	DeleteProject(id uint) error
	UpdateProject(id uint, project models.Projects) error
}

type ProjectsServiceImpl struct {
	DB *gorm.DB
}

func (s *ProjectsServiceImpl) CreateProject(project *models.Projects) error {
	if err := s.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func (s *ProjectsServiceImpl) GetProjectsByPage(page int) ([]models.Projects, error) {
	var projects []models.Projects
	if err := s.DB.Limit(10).Offset(page).Preload("Careers").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *ProjectsServiceImpl) GetProjectsByID(id uint) (*models.Projects, error) {
	var project models.Projects
	if err := s.DB.Where("id = ?", id).Preload("Careers").First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (s *ProjectsServiceImpl) UpdateProject(id uint, project models.Projects) error {
	if err := s.DB.Model(&project).Where("id = ?", id).Updates(&project).Error; err != nil {
		return err
	}
	return nil
}

func (s *ProjectsServiceImpl) DeleteProject(id uint) error {
	var project models.Projects
	if err := s.DB.Where("id = ?", id).Delete(&project).Error; err != nil {
		return err
	}
	return nil
}
