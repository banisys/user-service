package handlers

import (
	"net/http"
	"strconv"

	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/internal/services"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

type UserHandler struct {
	UserService services.UserService
}

func (h *UserHandler) Signup(context *gin.Context) {

	// fmt.Println("########################")

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

func (h *UserHandler) Login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = h.UserService.ValidateCredentials(&user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := h.UserService.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

}

func (h *UserHandler) Update(context *gin.Context) {

	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse user id."})
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedUser.ID = userId

	err = h.UserService.UpdateUser(&updatedUser)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})

}
