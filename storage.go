package main

import (
	"sync"
	"time"
)

type MemoryStorage struct {
    mu      sync.RWMutex
    readings []AirReading
}

func NewStorage() *MemoryStorage {
    return &MemoryStorage{readings: make([]AirReading, 0)}
}

func (s *MemoryStorage) Add(reading AirReading) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.readings = append(s.readings, reading)
}

func (s *MemoryStorage) GetNearest(ts time.Time) *AirReading {
    s.mu.RLock()
    defer s.mu.RUnlock()

    var closest *AirReading
    var smallestDiff time.Duration

    for _, r := range s.readings {
        diff := absDuration(r.Timestamp.Sub(ts))
        if closest == nil || diff < smallestDiff {
            closest = &r
            smallestDiff = diff
        }
    }
    return closest
}

func absDuration(d time.Duration) time.Duration {
    if d < 0 {
        return -d
    }
    return d
}
