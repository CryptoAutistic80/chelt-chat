package tests

import (
    "net/http"
    "testing"
)

func TestBackendIsRunning(t *testing.T) {
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        t.Fatalf("Failed to reach backend: %v", err)
    }
    if resp.StatusCode != 200 {
        t.Fatalf("Expected status 200, got %d", resp.StatusCode)
    }
} 