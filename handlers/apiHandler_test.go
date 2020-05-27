package handlers_test

import (
	"learnable-be/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexGets(t *testing.T) {
	t.Run("Test AppIndex endpoint", func(t *testing.T) {
		expected := `{"message":"Welcome to Learnable","status":200}`

		router := routes.GetRoutes()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("Test APIIndex endpoint", func(t *testing.T) {
		expected := `{"message":"Welcome to Learnable API","status":200}`

		router := routes.GetRoutes()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v0", nil)
		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expected, w.Body.String())
	})
}
