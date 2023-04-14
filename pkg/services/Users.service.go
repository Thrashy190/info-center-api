package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/utils"
	"gorm.io/gorm"
)

type UsersService interface {
	SignUpUser(user models.Users) (*models.Users, error)
	GetUsersByPage(page uint) ([]models.Users, error)
	GetUserByID(id uint) (*models.Users, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, tesis models.Users) error
}

type UsersServiceImpl struct {
	DB *gorm.DB
}

func (s *UsersServiceImpl) SignUpUser(user models.Users) (*models.Users, error) {
	var err error

	if user.Password, err = utils.HashPassword(user.Password); err != nil {
		return nil, err
	}

	println(user.Password)

	return &user, nil
}

func (s *UsersServiceImpl) GetUsersByPage(page int) ([]models.Users, error) {
	var users []models.Users
	if err := s.DB.Limit(10).Offset(page).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UsersServiceImpl) GetUserByID(id uint) (*models.Users, error) {
	var user models.Users
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersServiceImpl) UpdateUser(id uint, user models.Users) error {
	if err := s.DB.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UsersServiceImpl) DeleteUser(id uint) error {
	var user models.Users
	if err := s.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
