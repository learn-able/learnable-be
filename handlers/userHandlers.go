package handlers

import (
	"log"
	"net/http"
	"time"

	models "learnable-be/models"

	"github.com/gin-gonic/gin"
)

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
	// Create variable to hold the new user
	var user models.User

	// Grabs and assigns information retrieved from JSON body
	// checks returned error for any probems
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		panic(bindErr)
	}

	// stores incoming data in variables for struct assignment
	username := user.Username
	password := user.Password

	// Inserts in to next row of table with new user
	insertionError := models.DBConnect.Insert(&models.User{
		ID:        user.ID,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	// Validates and parses error, returning an internal server error
	if insertionError != nil {
		log.Printf("Could not insert user\nReason: %v\n", insertionError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	// Returns status ok and 201 if the use was successfully created
	// Also returns new user.
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created user",
		"data":    &user,
	})
}

func EditUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "EDIT ENDPOINT",
	})
}

func OneUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "SINGLE USER ENDPOINT",
	})
	// userID := mux.Vars(r)["id"]

	// for _, oneUser := range users {
	// 	if oneUser.ID == userID {
	// 		json.NewEncoder(w).Encode(oneUser)
	// 	}
	// }
}
