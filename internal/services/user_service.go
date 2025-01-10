package services

import (
	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Save(user)
}
