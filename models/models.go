package models

import "time"

type Reading struct {
	ID          int
	DeviceID    string
	Timestamp   time.Time
	Consumption float64
}
