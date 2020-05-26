package config

import (
	"log"
	"net/url"
	"os"

	models "learnable-be/models"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	// For HEROKU DEPLOYMENT
	parsedUrl, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	pgOptions := &pg.Options{
		User:     parsedUrl.User.Username(),
		Database: parsedUrl.Path[1:],
		Addr:     parsedUrl.Host,
	}

	if password, ok := parsedUrl.User.Password(); ok {
		pgOptions.Password = password
	}

	db := pg.Connect(pgOptions)
	// END HEROKU DEPLOYMENT

	// postgres setup parameters
	//   Use for local config ONLY
	// options := &pg.Options{
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Addr:     os.Getenv("DB_ADDRESS"),
	// 	Database: os.Getenv("DB_NAME"),
	// }

	// var db := pg.Connect(options)
	// End of local config block

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
