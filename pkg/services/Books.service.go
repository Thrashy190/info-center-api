package service

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type BooksService interface {
	GetBooksByPage(page uint) ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	CreateBook(*models.Book) error
	DeleteBook(id uint) error
	UpdateBook(id uint, book models.Book) error
}

type BooksServiceImpl struct {
	DB *gorm.DB
}

func (s *BooksServiceImpl) CreateBook(book *models.Book) error {
	if err := s.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (s *BooksServiceImpl) GetBooksByPage(page int) ([]models.Book, error) {
	var books []models.Book
	if err := s.DB.Limit(5).Offset(page).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BooksServiceImpl) GetBooksByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := s.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BooksServiceImpl) UpdateBook(id uint, book models.Book) error {
	if err := s.DB.Model(&book).Where("id = ?", id).Updates(&book).Error; err != nil {
		return err
	}
	return nil
}

func (s *BooksServiceImpl) DeleteBook(id uint) error {
	var book models.Book
	if err := s.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}
