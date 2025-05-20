package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	GoHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", resp.StatusCode)
	}
}
