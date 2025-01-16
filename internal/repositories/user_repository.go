package repositories

import (
	"github.com/banisys/user-service/internal/models"
)

type UserRepository interface {
	Save(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	Update(user *models.User) error
}
