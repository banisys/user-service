package services

import "github.com/banisys/user-service/internal/models"

type UserService interface {
	Create(user *models.User) error
	ValidateCredentials(user *models.User) error
	GenerateToken(email string, userId int64) (string, error)
}
