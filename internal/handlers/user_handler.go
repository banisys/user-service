package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/banisys/user-service/internal/models"
	"github.com/banisys/user-service/internal/services"
	"github.com/banisys/user-service/pkg/database"
	"github.com/banisys/user-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"context"

	pb "github.com/banisys/user-service/user_service_grpc"
)

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

type UserHandler struct {
	UserService services.UserService
}

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) Signup(_ context.Context, in *pb.UserServiceReq) (*pb.UserServiceRes, error) {

	var user models.User

	user.Name = in.GetName()
	user.Email = in.GetEmail()
	user.Password = in.GetPassword()

	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	stmt, err := database.DB().Prepare(query)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot query prepare")

	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return &pb.UserServiceRes{Message: "err"}, err
	}

	result, err := stmt.Exec(user.Name, user.Email, hashedPassword)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot query execute")
	}

	if err != nil {
		log.Error().Err(err)
	}

	fmt.Println(result)

	return &pb.UserServiceRes{Message: "User created successfully"}, nil

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
