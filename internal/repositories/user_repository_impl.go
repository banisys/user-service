package repositories

import (
	"database/sql"
	"errors"

	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/pkg/utils"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
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

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := r.DB.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
