package services

import (
	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/internal/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (u *UserServiceImpl) Create(user *models.User) error {
	if err := u.UserRepository.Save(user); err != nil {
		return err
	}
	return nil
}
