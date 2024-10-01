package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from backend! You requested: %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Backend server started on :8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal("ListenAndServe error:", err)
    }
}
