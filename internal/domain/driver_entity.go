// Package domain represents the domain layer of the application.
package domain

import "time"

// DriverPosition represents a delivery driver entity.
type DriverPosition struct {
	ID        uint64    `json:"id,omitempty"`
	DriverID  uint64    `json:"driverId"`
	Location  Point     `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}

// Point represents a geographic point with latitude and longitude coordinates.
type Point struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}
