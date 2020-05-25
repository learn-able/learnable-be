package handlers

import "github.com/gin-gonic/gin"

type CreatePlaylistInput struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	DueDate string `json:"due_date"`
}

type ReturnedPlaylist struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	DueDate string `json:"due_date"`
}

func CreatePlaylist(c *gin.Context) {
	// TODO
}

func UserPlaylists(c *gin.Context) {
	// TODO
}

func ShowPlaylist(c *gin.Context) {
	// TODO
}
