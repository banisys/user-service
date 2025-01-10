package routes

import (
	"github.com/banisys/user-service/internal/handlers"
	"github.com/gin-gonic/gin"

	"github.com/banisys/user-service/internal/repositories"
	"github.com/banisys/user-service/internal/services"
	"github.com/banisys/user-service/pkg/database"
)

func RegisterRoutes(route *gin.Engine) {

	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	route.POST("/signup", userHandler.Signup)

}
