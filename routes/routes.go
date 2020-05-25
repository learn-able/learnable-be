package routes

import (
	"github.com/gin-gonic/gin"
	handlers "learnable-be/handlers"
)

func GetRoutes(router *gin.Engine) {
	// User related routes
	router.GET("/api/v0", handlers.IndexHandler)
	router.GET("api/v0/users", handlers.AllUsersHandler)
	router.GET("/api/v0/users/:id", handlers.OneUserHandler)
	router.POST("/api/v0/users", handlers.CreateUserHandler)
}
