package routes

import (
	"GoAuth/controllers"
	"GoAuth/middlewares"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(server *gin.Engine) {

	public := server.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := server.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/users", controllers.AllUsers)
	protected.GET("/user", controllers.CurrentUser)

}
