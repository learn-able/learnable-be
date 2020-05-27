package handlers

import (
	"learnable-be/models"
	"log"
	"net/http"
	"strconv"

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
	if err := models.PlaylistConnect.Insert(&playlist); err != nil {
		panic(err)
	}

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
	var input CreatePlaylistInput

	if bindErr := c.BindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	var playlists []models.Playlist

	err := models.PlaylistConnect.
		Model(&playlists).
		Where("user_id = ?", input.UserID).
		Select()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All playlists by User",
		"data":    &playlists,
	})
}

func PlaylistsByStatus(c *gin.Context) {
	var input CreatePlaylistInput

	if bindErr := c.BindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	var playlists []models.Playlist

	err := models.PlaylistConnect.
		Model(&playlists).
		Where("user_id = ? AND status = ?", input.UserID, input.Status).
		Select()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All playlists by Status",
		"data":    &playlists,
	})
}

func ShowPlaylist(c *gin.Context) {
	db := models.PlaylistConnect
	playlistID, _ := strconv.Atoi(c.Param("id"))
	playlist := &models.Playlist{ID: playlistID}

	err := db.Select(playlist)
	if err != nil {
		log.Printf("Error retrieving Playlist from database\nReason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Playlist not found",
		})
		return
	}

	foundPlaylist := ReturnedPlaylist{
		ID:            playlist.ID,
		UserID:        playlist.UserID,
		Title:         playlist.Title,
		Status:        playlist.Status,
		DueDate:       playlist.DueDate,
		PlaylistItems: playlist.PlaylistItems,
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Playlist found",
		"data":    foundPlaylist,
	})
}
