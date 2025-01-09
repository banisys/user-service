package main

import (
	"github.com/banisys/user-service/internal/routes"
	"github.com/banisys/user-service/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.CreateTables()

	route := gin.Default()

	routes.RegisterRoutes(route)

	route.Run()
}
