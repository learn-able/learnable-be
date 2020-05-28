package handlers

import (
	"learnable-be/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistInput struct {
	ID      int    `json:"id"`
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
		PlaylistItems: []*models.PlaylistItem{},
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created playlist",
		"data":    newPlaylist,
	})
}

func UserPlaylists(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))

	var playlists []models.Playlist

	err := models.PlaylistConnect.
		Model(&playlists).
		Where("user_id = ?", userID).
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
	q := c.Request.URL.Query()
	status := q.Get("status")
	userID := q.Get("user_id")

	var playlists []models.Playlist

	err := models.PlaylistConnect.
		Model(&playlists).
		Where("user_id = ? AND status = ?", userID, status).
		Select()

	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All playlists by Status",
		"data":    playlists,
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
		PlaylistItems: []*models.PlaylistItem{},
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Playlist found",
		"data":    foundPlaylist,
	})
}

func UpdatePlaylist(c *gin.Context) {
	var input CreatePlaylistInput

	if bindErr := c.BindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	playlistID, _ := strconv.Atoi(c.Param("id"))

	playlist := models.Playlist{
		UserID:  input.UserID,
		Title:   input.Title,
		Status:  input.Status,
		DueDate: input.DueDate,
	}

	_, err := models.PlaylistConnect.Model(&playlist).
		Where("id = ?", playlistID).
		Update(&playlist)

	if err != nil {
		panic(err)
	}

	foundPlaylist := ReturnedPlaylist{
		ID:            playlistID,
		UserID:        playlist.UserID,
		Title:         playlist.Title,
		Status:        playlist.Status,
		DueDate:       playlist.DueDate,
		PlaylistItems: []*models.PlaylistItem{},
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Playlist updated",
		"data":    foundPlaylist,
	})
}
