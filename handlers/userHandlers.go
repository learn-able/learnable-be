package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Defined user Struct for use in json responses
type User struct {
	ID       	int 			`json:"id" pg:"pk_id"`
	Username 	string 		`json:"username" pg:",unique"`
	Password	string		`json:"password"`
	Points   	int    		`json:"points"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Initialize users table for use in API
func CreateUserTable(db *pg.DB) error {
	// Only created this table if it doesn't already exist
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	// logs any errors for return later related to the creation of this users table
	err := db.CreateTable(&User{}, options)

	if err != nil {
		log.Printf("Cannot create User table.\nReason: %v\n", err)
		return err
	}

	log.Printf("User table created successfully.")
	return nil
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
