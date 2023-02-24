package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RequestService interface {
	GetRequestsByPage(page uint) ([]models.Request, error)
	GetRequestByID(id uint) (*models.Request, error)
	CreateRequest(*models.Request) error
	UpdateRequest(id uint) error
}

type RequestServiceImpl struct {
	DB *gorm.DB
}

func (s *RequestServiceImpl) CreateRequest(request *models.Request) error {
	if err := s.DB.Create(request).Error; err != nil {
		return err
	}
	return nil
}

func (s *RequestServiceImpl) GetRequestByID(id uint) (*models.Request, error) {
	var request models.Request
	if err := s.DB.Where("id = ?", id).Preload(clause.Associations).First(&request).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (s *RequestServiceImpl) GetRequestsByPage(page int) ([]models.Request, error) {
	var requests []models.Request
	if err := s.DB.Limit(5).Offset(page).Preload(clause.Associations).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (s *RequestServiceImpl) UpdateRequest(id uint) error {

	if err := s.DB.Model(models.Request{}).Where("id = ?", id).Update("state_request", true).Error; err != nil {
		return err
	}
	return nil
}
