package models

type User struct {
	ID       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
