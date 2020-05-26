package handlers

import (
	"encoding/json"
	"learnable-be/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAPIIndex(t *testing.T) {
	body := gin.H{
		"status":  200,
		"message": "Welcome To Learnable API",
	}

	router := routes.GetRoutes()

	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["status"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["status"], value)
}
