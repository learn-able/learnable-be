package routes

import (
	handlers "learnable-be/handlers"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) {
	// User related routes
	router.GET("/api/v0", handlers.IndexHandler)
	router.GET("/api/v0/users/:id", handlers.OneUserHandler)
	router.POST("/api/v0/users", handlers.CreateUserHandler)
}
