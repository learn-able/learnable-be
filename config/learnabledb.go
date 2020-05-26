package config

import (
	"log"
	"os"
	"strings"

	models "learnable-be/models"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	url := os.Getenv("DATABASE_URL")
	url = strings.TrimPrefix(url, "postgres://")

	dbNameStartsAt := strings.LastIndex(url, "/") + 1
	database := url[dbNameStartsAt:]
	url = url[:dbNameStartsAt-1]

	authAndHost := strings.Split(url, "@")
	auth := strings.Split(authAndHost[0], ":")
	username := auth[0]
	password := auth[1]
	hostAndPort := authAndHost[1]

	db := pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: database,
		Addr:     hostAndPort,
	})
	// postgres setup parameters
	// options := &pg.Options{
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Addr:     os.Getenv("DB_ADDRESS"),
	// 	Database: os.Getenv("DB_NAME"),
	// }

	// var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Printf("Could not connect to Learnable Database")
		os.Exit(100)
	}

	log.Printf("Connected to Learnable Database")

	createTables(db)

	return db
}

func createTables(db *pg.DB) {
	if userTblErr := models.CreateUserTable(db); userTblErr != nil {
		panic(userTblErr)
	}

	models.InitiateUserDB(db)

	if playlistTblErr := models.CreatePlaylistTable(db); playlistTblErr != nil {
		panic(playlistTblErr)
	}

	models.InitiatePlaylistDB(db)
}
