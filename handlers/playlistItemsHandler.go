package handlers

import (
	"learnable-be/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePlaylistItemInput struct {
	PlaylistID  int       `json:"playlist_id"`
  Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	IsComplete  bool      `json:"is_complete"`
}

type ReturnedPlaylistItem struct {
	ID            int       `json:"id" pg:"pk_id"`
	PlaylistID    int       `json:"playlist_id"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Description   string    `json:"description,omitempty"`
	URL           string    `json:"url"`
  IsComplete    bool      `json:"is_complete"`
}

func CreatePlaylistItem(c *gin.Context) {
	var input CreatePlaylistItemInput

	if bindErr := c.BindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	playlistItem := models.PlaylistItem{
    PlaylistID:  input.PlaylistID,
    Name:        input.Name,
		Category:    input.Category,
  	Description: input.Description,
  	URL:         input.URL,
  	IsComplete:  false,
	}
	if err := models.PlaylistItemConnect.Insert(&playlistItem); err != nil {
		panic(err)
	}

	newPlaylistItem := ReturnedPlaylistItem {
    ID:             playlistItem.ID,
  	PlaylistID:     playlistItem.PlaylistID,
  	Name:           playlistItem.Name,
		Category:       playlistItem.Category,
  	Description:    playlistItem.Description,
  	URL:            playlistItem.URL,
    IsComplete:     playlistItem.IsComplete,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Playlist found",
		"data":    foundPlaylist,
	})
}

func ShowPlaylistItem(c *gin.Context) {
	db := models.PlaylistItemConnect
	playlistItemID, _ := strconv.Atoi(c.Param("id"))
	playlistItem := &models.PlaylistItem{ID: playlistItemID}

	err := db.Select(playlistItem)
	if err != nil {
		log.Printf("Error retrieving Playlist Item from database\nReason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Playlist Item not found",
		})
		return
	}

	foundPlaylistItem := ReturnedPlaylistItem {
    ID:             playlistItem.ID,
  	PlaylistID:     playlistItem.PlaylistID,
  	Name:           playlistItem.Name,
		Category:       playlistItem.Category,
  	Description:    playlistItem.Description,
  	URL:            playlistItem.URL,
    IsComplete:     playlistItem.IsComplete,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Playlist found",
		"data":    foundPlaylistItem,
	})
}

func PlaylistItems(c *gin.Context) {
	playlistID, _ := strconv.Atoi(c.Param("playlist_id"))
	var playlistItems []models.PlaylistItem

	err := models.PlaylistItemConnect.
		Model(&playlistItems).
		Where("playlist_id = ?", playlistID).
		Select()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All playlists items by playlist",
		"data":    &playlistItems,
	})
}
