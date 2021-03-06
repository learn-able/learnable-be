package models

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Playlist struct {
	ID            int             `json:"id" pg:"pk_id"`
	UserID        int             `json:"user_id"`
	Title         string          `json:"title"`
	Status        string          `json:"status"`
	DueDate       string          `json:"due_date"`
	PlaylistItems []*PlaylistItem `json:"playlist_items"`
	IsFavorite    bool            `json:"is_favorite"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

var PlaylistConnect *pg.DB

func CreatePlaylistTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	if err := db.CreateTable(&Playlist{}, options); err != nil {
		log.Printf("Cannot create Playlist table.\nReason: %v\n", err)
		return err
	}

	log.Printf("Playlist table created successfully.")
	return nil
}

func InitiatePlaylistDB(db *pg.DB) {
	PlaylistConnect = db
}
