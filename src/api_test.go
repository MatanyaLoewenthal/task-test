package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetAlbums(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/albums", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Test that the http status code is 200
	assert.Equal(t, 200, w.Code)

	// Test that the response body is correct
}
