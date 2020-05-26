package handlers

import (
	"learnable-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistInput struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	DueDate string `json:"due_date"`
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
	var input CreatePlaylistInput

	if bindErr := c.BindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	playlist := models.Playlist{
		UserID:  input.UserID,
		Title:   input.Title,
		Status:  "valid",
		DueDate: input.DueDate,
	}
	models.PlaylistConnect.Insert(&playlist)

	newPlaylist := ReturnedPlaylist{
		ID:            playlist.ID,
		UserID:        playlist.UserID,
		Title:         playlist.Title,
		Status:        playlist.Status,
		DueDate:       playlist.DueDate,
		PlaylistItems: playlist.PlaylistItems,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created playlist",
		"data":    newPlaylist,
	})
}

func UserPlaylists(c *gin.Context) {
	// TODO
}

func ShowPlaylist(c *gin.Context) {
	// TODO
}
