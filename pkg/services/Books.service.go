package service

import (
	"math"

	"github.com/Thrashy190/info-center-api/pkg/models"
	"gorm.io/gorm"
)

type BooksService interface {
	GetBooksByPage(page uint) (*models.PaginationDataBooks, error)
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

func (s *BooksServiceImpl) GetBooksByPage(page int) (*models.PaginationDataBooks, error) {

	perPage := 4

	//Count total books
	var totalRows int64
	if err := s.DB.Model(&models.Book{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}
	totalPages := math.Ceil(float64(totalRows) / float64(perPage))

	offset := (page - 1) * perPage

	//Get pagination data
	var books []models.Book
	if err := s.DB.Limit(4).Offset(offset).Find(&books).Error; err != nil {
		return nil, err
	}

	var paginationData = models.PaginationDataBooks{
		Books:        &books,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   int(totalPages),
	}

	return &paginationData, nil
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
