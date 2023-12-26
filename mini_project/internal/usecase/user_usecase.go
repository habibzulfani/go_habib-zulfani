package usecase

import (
	"errors"
	"mini_project/internal/domain"
)

type UserUseCase interface {
	Register(user *domain.User) error
	Login(username, password string) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uint) error
}

type userUseCase struct {
	userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (uc *userUseCase) Register(user *domain.User) error {
	return uc.userRepository.Create(user)
}

func (uc *userUseCase) Login(username, password string) (*domain.User, error) {
	user, err := uc.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	// Check if the password matches (you may want to use a secure password hashing library)
	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}

func (uc *userUseCase) GetUserByID(id uint) (*domain.User, error) {
	return uc.userRepository.GetByID(id)
}

func (uc *userUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.userRepository.GetAll()
}

func (uc *userUseCase) UpdateUser(user *domain.User) error {
	return uc.userRepository.Update(user)
}

func (uc *userUseCase) DeleteUser(id uint) error {
	return uc.userRepository.Delete(id)
}
