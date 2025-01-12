package handlers

import (
	"net/http"

	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (h *UserHandler) Signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = h.UserService.Create(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
