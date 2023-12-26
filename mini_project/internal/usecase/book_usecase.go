package usecase

import (
	"mini_project/internal/domain"
)

// BookUseCase provides use cases for books
type BookUseCase interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	CreateBook(book *domain.Book) error
	DeleteBook(id uint) error
}

type bookUseCase struct {
	bookRepository domain.BookRepository
}

// NewBookUseCase creates a new instance of BookUseCase
func NewBookUseCase(bookRepository domain.BookRepository) BookUseCase {
	return &bookUseCase{bookRepository: bookRepository}
}

func (uc *bookUseCase) GetAllBooks() ([]domain.Book, error) {
	return uc.bookRepository.GetAll()
}

func (uc *bookUseCase) GetBookByID(id uint) (*domain.Book, error) {
	return uc.bookRepository.GetByID(id)
}

func (uc *bookUseCase) CreateBook(book *domain.Book) error {
	return uc.bookRepository.Create(book)
}

func (uc *bookUseCase) DeleteBook(id uint) error {
	return uc.bookRepository.Delete(id)
}

func (uc *bookUseCase) GetUserBooks(userID uint) ([]domain.Book, error) {
	// Implement logic to retrieve books based on user's role or other criteria

	// For simplicity, let's assume that users can only see books associated with their user ID
	return uc.bookRepository.GetByUserID(userID)
}