package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chelt Chat Backend is Running")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Backend listening on port 8080")
    http.ListenAndServe(":8080", nil)
} 