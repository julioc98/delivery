package app

import "github.com/julioc98/delivery/internal/domain"

// DeliveryRepository represents a repository for delivery drivers.
type DeliveryRepository interface {
	// SaveDriverPosition saves a driver position.
	SaveDriverPosition(driverID uint64, latitude, longitude float64) (int64, error)
	// FindDriverPosition finds a driver position.
	FindDriverPosition(driverID uint64) (domain.DriverPosition, error)
	// HistoryDriverPosition finds a driver position history.
	HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error)
	// GetDriversNearby finds drivers nearby.
	GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error)
}

// Publisher represents a publisher.
type Publisher interface {
	Publish(subject string, data []byte) error
}
