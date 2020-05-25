package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Defined user Struct for use in json responses
type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Points   int    `json:"points"`
}

// preps an array of all of the users stored in memory
type allUsers []user

// Creates a 'dummy' user for debugging and endpoint handling
var users = allUsers{
	{
		ID:       "1",
		Username: "Learnable_User",
		Points:   1000,
	},
}

func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "CREATE ENDPOINT",
		})
		return
	// var newUser user
	// reqBody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, "Enter data for your new user")
	// }

	// json.Unmarshal(reqBody, &newUser)

	// users = append(users, newUser)
	// w.WriteHeader(http.StatusCreated)

	// json.NewEncoder(w).Encode(newUser)
}

func OneUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "SINGLE USER ENDPOINT",
		})
		return
	// userID := mux.Vars(r)["id"]

	// for _, oneUser := range users {
	// 	if oneUser.ID == userID {
	// 		json.NewEncoder(w).Encode(oneUser)
	// 	}
	// }
}

func AllUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "ALL USER ENDPOINT",
		})
		return
	// json.NewEncoder(w).Encode(users)
}
