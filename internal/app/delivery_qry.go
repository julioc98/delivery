package app

import (
	"github.com/julioc98/delivery/internal/domain"
)

// DeliveryQry represents a NATS message producer.
type DeliveryQry struct {
	repo DeliveryFinder
}

// NewDeliveryQry creates a new DeliveryQry.
func NewDeliveryQry(r DeliveryFinder) *DeliveryQry {
	return &DeliveryQry{repo: r}
}

// FindDriverPosition finds a driver position.
func (q *DeliveryQry) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {
	return q.repo.FindDriverPosition(driverID)
}

// HistoryDriverPosition finds a driver position history.
func (q *DeliveryQry) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	return q.repo.HistoryDriverPosition(driverID)
}

// GetDriversNearby finds drivers nearby.
func (q *DeliveryQry) GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error) {
	return q.repo.GetDriversNearby(latitude, longitude, radius)
}
