package producer

import (
	"math/rand"

	"time"
)

type DeviceReading struct {
	DeviceID    string
	Timestamp   time.Time
	Consumption float64
}

func SimulateReading(deviceID string) DeviceReading {
	rand.Seed(time.Now().UnixNano())
	return DeviceReading{
		DeviceID:    deviceID,
		Timestamp:   time.Now(),
		Consumption: 0.5 + rand.Float64()*5,
	}
}
