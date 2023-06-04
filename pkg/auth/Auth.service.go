package auth

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type AuthenticationService interface {
	Login(*models.Auth) *models.Users
	SignUp(*models.Users) error
	SignOut(*models.Users) error
}

type AuthenticationServiceImpl struct {
	DB *gorm.DB
}

func (s *AuthenticationServiceImpl) SignUp(user *models.Users) error {
	if err := s.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *AuthenticationServiceImpl) Login(auth *models.Auth) *models.Users {
	var user models.Users
	if err := s.DB.Where("email = ?", auth.Email).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
