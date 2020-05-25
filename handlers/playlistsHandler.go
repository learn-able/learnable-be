package handlers

import (
	"learnable-be/models"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistInput struct {
	UserID        int                    `json:"user_id"`
	Title         string                 `json:"title"`
	Status        string                 `json:"status"`
	DueDate       string                 `json:"due_date"`
	PlaylistItems []*models.PlaylistItem `json:"items"`
}

type ReturnedPlaylist struct {
	ID            int                    `json:"id"`
	UserID        int                    `json:"user_id"`
	Title         string                 `json:"title"`
	Status        string                 `json:"status"`
	DueDate       string                 `json:"due_date"`
	PlaylistItems []*models.PlaylistItem `json:"playlist_items"`
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
