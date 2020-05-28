package models

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type PlaylistItem struct {
	ID          int       `json:"id" pg:"pk_id"`
	PlaylistID  int       `json:"playlist_id" pg:"fk:Playlist"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	IsComplete  bool      `json:"is_complete"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var PlaylistItemConnect *pg.DB

func CreatePlaylistItemTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	if err := db.CreateTable(&PlaylistItem{}, options); err != nil {
		log.Printf("Cannot create PlaylistItems table.\nResponse: %v\n", err)
		return err
	}

	log.Printf("PlaylistItems table created successfully.")
	return nil
}

func InitiatePlaylistItemDB(db *pg.DB) {
	PlaylistItemConnect = db
}
