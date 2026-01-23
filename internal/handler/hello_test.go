package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint_ReturnsCorrectResponse(t *testing.T) {
	// Arrange
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// Act
	handler.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!\n", rec.Body.String())
}

func TestHelloEndpoint_MethodNotAllowed(t *testing.T) {
	// Arrange
	req := httptest.NewRequest("POST", "/hello", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// Act
	handler.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Contains(t, rec.Body.String(), "405 Method Not Allowed")
}