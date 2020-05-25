package config

import (
	"log"
	"os"

	handlers "learnable-be/handlers"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

func Connect() *pg.DB {
	// Add environment variables for use in db auth
	envErr := godotenv.Load("application.env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
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

	tblErr := handlers.CreateUserTable(db)
	if tblErr != nil {
		panic(tblErr)
	}

	return db
}
