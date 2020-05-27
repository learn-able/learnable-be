package handlers_test

import (
	"bytes"
	"learnable-be/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUserActions(t *testing.T) {
	t.Run("Test Get User by ID endpoint", func(t *testing.T) {
		router := routes.GetRoutes()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v0/users/10", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})

	t.Run("Test Create a User", func(t *testing.T) {
		newUser := `{"username": "learnable_1","password": "123456"}`
		router := routes.GetRoutes()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/api/v0/users", bytes.NewBufferString(newUser))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
}
