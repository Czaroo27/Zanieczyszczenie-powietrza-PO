package main

import "time"

type AirReading struct {
    Timestamp   time.Time `json:"timestamp"`
    PM25        float64   `json:"pm2_5"`
    PM10        float64   `json:"pm10"`
    CO          float64   `json:"carbon_monoxide"`
}
