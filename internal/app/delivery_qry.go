// Package mdg represents the message driven gateway.
package app

import (
	"github.com/julioc98/delivery/internal/domain"
)

// DeliveryQry represents a NATS message producer.
type DeliveryQry struct {
	publisher Publisher
}

// NewDeliveryQry creates a new DeliveryQry.
func NewDeliveryQry(publisher Publisher) *DeliveryQry {
	return &DeliveryQry{publisher: publisher}
}

// FindDriverPosition retrieves a driver's position from either the Redis cache or the database.
func (p *DeliveryQry) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {

	return domain.DriverPosition{}, nil
}

// HistoryDriverPosition finds a driver position history.
func (p *DeliveryQry) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	return nil, nil
}

// GetDriversNearby finds drivers nearby.
func (p *DeliveryQry) GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error) {
	return nil, nil
}
