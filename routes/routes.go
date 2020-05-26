package routes

import (
	handlers "learnable-be/handlers"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) {
	// API Index
	router.GET("/api/v0", handlers.Index)

	// User related routes
	router.POST("/api/v0/users", handlers.CreateUser)
	router.GET("/api/v0/users/:id", handlers.ShowUser)

	// Playlist related routes
	router.POST("/api/v0/playlists", handlers.CreatePlaylist)
	router.GET("/api/v0/playlists", handlers.UserPlaylists)
	router.GET("/api/v0/playlists/:id", handlers.ShowPlaylist)
}
