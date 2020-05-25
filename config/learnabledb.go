package config

import (
	"log"
	"os"

	handlers "learnable-be/handlers"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	// postgres setup options
	options := &pg.Options{
		User: "admin",
		Password: "postgres",
		Addr: "localhost:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Printf("Could not connect to Learnable Database")
		os.Exit(100)
	}

	log.Printf("Connected to Learnable Database")

	handlers.CreateUserTable(db)
	return db
}