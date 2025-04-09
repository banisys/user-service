package routes

import (
	"github.com/banisys/user-service/internal/handlers"
	"github.com/banisys/user-service/internal/middlewares"
	"github.com/banisys/user-service/internal/repositories"
	"github.com/banisys/user-service/internal/services"
	"github.com/banisys/user-service/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	userRepository := repositories.NewUserRepository(database.DB())
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// route.POST("/signup", userHandler.Signup)
	route.POST("/login", userHandler.Login)

	auth := route
	auth.Use(middlewares.Authenticate)

	auth.PUT("/users/:id", userHandler.Update)

}
