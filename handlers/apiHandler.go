package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandlers returns a json response for the index location of the API with welcome and status.
func APIIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Welcome to Learnable API",
		})
}

func AppIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Welcome to Learnable",
	})
}
