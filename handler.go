package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type AirHandler struct {
    storage *MemoryStorage
}

func NewAirHandler(storage *MemoryStorage) *AirHandler {
    return &AirHandler{storage}
}

func (h *AirHandler) PostReading(w http.ResponseWriter, r *http.Request) {
    var reading AirReading
    if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
        http.Error(w, "Nieprawidłowe dane wejściowe", http.StatusBadRequest)
        return
    }

    // Walidacja
    if reading.PM25 < 0 || reading.PM10 < 0 || reading.CO < 0 {
        http.Error(w, "Dane muszą być większe od zera", http.StatusBadRequest)
        return
    }

    h.storage.Add(reading)
    w.WriteHeader(http.StatusCreated)
}

func (h *AirHandler) GetNearestReading(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    tsStr := query.Get("timestamp")
    if tsStr == "" {
        http.Error(w, "Brak parametru timestamp", http.StatusBadRequest)
        return
    }

    ts, err := time.Parse(time.RFC3339, tsStr)
    if err != nil {
        http.Error(w, "Nieprawidłowy format czasu (wymagany RFC3339)", http.StatusBadRequest)
        return
    }

    result := h.storage.GetNearest(ts)
    if result == nil {
        http.Error(w, "Brak odczytów", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
