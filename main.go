package main

import (
	"GoAuth/models"
	"GoAuth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	server := gin.Default()

	routes.InitializeRoutes(server)

	server.Run(":8080")
}
