package routes

import (
	handlers "learnable-be/handlers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()

	// Application
	router.GET("/", handlers.AppIndex)

	// API Index
	router.GET("/api/v0", handlers.APIIndex)

	// User related routes
	router.POST("/api/v0/users", handlers.CreateUser)
	router.GET("/api/v0/users/:id", handlers.ShowUser)

	// Playlist related routes
	router.POST("/api/v0/playlists", handlers.CreatePlaylist)
	router.GET("/api/v0/playlists", handlers.UserPlaylists)
	router.GET("/api/v0/playlists/:id", handlers.ShowPlaylist)

	return router
}
