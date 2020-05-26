package models

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Defined user Struct for use in json responses
type User struct {
	ID        int       `json:"id" pg:"pk_id"`
	Username  string    `json:"username" pg:"unique"`
	Password  string    `json:"password"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// instance variable used for this models connection to the db
var UserConnect *pg.DB

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

func InitiateUserDB(db *pg.DB) {
	UserConnect = db
}
