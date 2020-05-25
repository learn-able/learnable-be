package config

import (
	"log"
	"os"

	models "learnable-be/models"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	// postgres setup parameters
	options := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDRESS"),
		Database: os.Getenv("DB_NAME"),
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Printf("Could not connect to Learnable Database")
		os.Exit(100)
	}

	log.Printf("Connected to Learnable Database")

	tblErr := models.CreateUserTable(db)
	if tblErr != nil {
		panic(tblErr)
	}

	models.InitiateDB(db)

	return db
}
