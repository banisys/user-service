package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/internal/repositories"
	"github.com/banisys/user-service/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (u *UserServiceImpl) Create(user *models.User) error {
	fmt.Println("888888")
	if err := u.UserRepository.Save(user); err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) ValidateCredentials(user *models.User) error {

	retrievedUser, err := u.UserRepository.GetUserByEmail(user.Email)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedUser.Password)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func (u *UserServiceImpl) GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	if err := u.UserRepository.Update(user); err != nil {
		return err
	}
	return nil
}
