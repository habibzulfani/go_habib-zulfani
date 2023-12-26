package repository

import (
	"mini_project/internal/domain"
	"time"

	"gorm.io/gorm"
)

// BookRepository provides CRUD operations for books
type BookRepository interface {
	GetAll() ([]domain.Book, error)
	GetByID(id uint) (*domain.Book, error)
	Create(book *domain.Book) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]domain.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

// NewBookRepository creates a new instance of BookRepository
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*domain.Book, error) {
	var book domain.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Create(book *domain.Book) error {
	book.Published = time.Now()
	return r.db.Create(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Book{}, id).Error
}

func (r *bookRepository) GetByUserID(userID uint) ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Where("user_id = ?", userID).Find(&books).Error
	return books, err
}
