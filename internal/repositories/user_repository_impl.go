package repositories

import (
	"github.com/banisys/user-service/internal/models"

	"database/sql"

	"github.com/banisys/user-service/pkg/utils"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Save(user *models.User) error {

	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = userId

	return nil
}
