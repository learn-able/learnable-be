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

func AllUsersHandler(c *gin.Context) {
	// Creates an array to hold all users
	var users []models.User

	// Attempts a connection to the database to retrieve all users
	//   returns an Internal SErver Error if a problem occurs
	err := models.DBConnect.Model(&users).Select()
	if err != nil {
		log.Printf("Could not get all Users\nReason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	// Returns a JSON response for the user that is created
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Learnable Users",
		"data":    users,
	})
}

func CreateUserHandler(c *gin.Context) {
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
	models.DBConnect.Insert(&user)
	// Returns status ok and 201 if the use was successfully created
	// Also returns new user.
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created user",
		"data":    user,
	})
}

func OneUserHandler(c *gin.Context) {
	db := models.DBConnect
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

	// Returns status ok for user if all goes well.
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User found",
		"data":    user,
	})
}
