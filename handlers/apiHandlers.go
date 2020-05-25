package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandlers returns a json response for the index location of the API with welcome and status.
func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Learnable API",
		})
		return
}