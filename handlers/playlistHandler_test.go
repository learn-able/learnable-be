package handlers_test

import (
	"learnable-be/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistActions(t *testing.T) {
	t.Run("Test Get playlists by user ID endpoint", func(t *testing.T) {
		router := routes.GetRoutes()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v0/user/1/playlists", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})

	t.Run("Test can get playlist by ID", func(t *testing.T) {
		router := routes.GetRoutes()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v0/playlists/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
}
