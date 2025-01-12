package services

import "github.com/banisys/user-service/internal/models"

type UserService interface {
	Create(user *models.User) error
}
