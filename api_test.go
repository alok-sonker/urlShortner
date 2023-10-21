package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	mockResponse := `Welcome to short url service`
	server := initServer()
	server.GET("/", healthCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func BenchmarkHealthCheck(b *testing.B) {
	mockResponse := `Welcome to short url service`
	server := initServer()
	server.GET("/", healthCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(b, mockResponse, string(responseData))
	assert.Equal(b, http.StatusOK, w.Code)
}

func TestCreateURL(t *testing.T) {
	mockResponse := `https://url.com/1000000`
	server := initServer()
	server.GET("/", createURL)
	jsonBody := []byte(`{"url":"aloksonker.me"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest("GET", "/", bodyReader)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
