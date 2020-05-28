package routes

import (
	handlers "learnable-be/handlers"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// Application
	router.GET("/", handlers.AppIndex)

	// API Index
	router.GET("/api/v0", handlers.APIIndex)

	// User related routes
	router.POST("/api/v0/users", handlers.CreateUser)
	router.GET("/api/v0/users/:id", handlers.ShowUser)

	// Playlist related routes
	router.POST("/api/v0/playlists", handlers.CreatePlaylist)
	router.GET("/api/v0/user-playlists/:user_id", handlers.UserPlaylists)
	router.GET("/api/v0/playlists/:id", handlers.ShowPlaylist)
	router.PUT("/api/v0/playlists/:id", handlers.UpdatePlaylist)
	router.GET("/api/v0/playlists-status", handlers.PlaylistsByStatus)

	// PlaylistItem routes
	router.POST("/api/v0/playlist_items", handlers.CreatePlaylistItem)
	router.GET("/api/v0/playlist_items/:id", handlers.ShowPlaylistItem)
	router.GET("/api/v0/pl_playlist_items/:playlist_id", handlers.PlaylistItems)

	return router
}
