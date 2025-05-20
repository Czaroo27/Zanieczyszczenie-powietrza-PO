package main

import (
	"log"
	"net/http"
)

func main() {
    storage := NewStorage()
    handler := NewAirHandler(storage)

    http.HandleFunc("/api/reading", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            handler.PostReading(w, r)
        } else {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/api/nearest", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            handler.GetNearestReading(w, r)
        } else {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Serwer uruchomiony na porcie 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
