package handlers

import (
	"log"
	"net/http"
	"strconv"

	models "learnable-be/models"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReturnedUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Points   int    `json:"points"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	// Grabs and assigns information retrieved from JSON body
	// checks returned error for any probems
	if bindErr := c.BindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	// Creates user
	user := models.User{Username: input.Username, Password: input.Password}
	models.UserConnect.Insert(&user)

	// Returns status ok and 201 if the use was successfully created
	// Also returns new user.
	newUser := ReturnedUser{
		ID:       user.ID,
		Username: user.Username,
		Points:   user.Points,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created user",
		"data":    newUser,
	})
}

func ShowUser(c *gin.Context) {
	db := models.UserConnect
	userID, _ := strconv.Atoi(c.Param("id")) // Grab id from URI params and converts to int
	user := &models.User{ID: userID}         // Retrieve user from db through model

	// // Verify errors on not found by selection
	err := db.Select(user)
	if err != nil {
		log.Printf("Error retrieving user from database\nReason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	foundUser := ReturnedUser{
		ID:       user.ID,
		Username: user.Username,
		Points:   user.Points,
	}
	// Returns status ok for user if all goes well.
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User found",
		"data":    foundUser,
	})
}
