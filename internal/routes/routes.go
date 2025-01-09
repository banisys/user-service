package routes

import (
	"github.com/banisys/user-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	route.POST("/signup", handlers.Signup)

}
