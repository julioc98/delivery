// Package app implements the application layer.
package app

import "github.com/julioc98/delivery/internal/domain"

// DeliveryRepository represents a repository for delivery drivers.
type DeliveryRepository interface {
	// SaveDriverPosition saves a driver position.
	SaveDriverPosition(driverID uint64, latitude, longitude float64) error
	// FindDriverPosition finds a driver position.
	FindDriverPosition(driverID uint64) (domain.DriverPosition, error)
	// HistoryDriverPosition finds a driver position history.
	HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error)
}

// DeliveryUseCase represents a use case for delivery drivers.
type DeliveryUseCase struct {
	repo DeliveryRepository
}

// NewDeliveryUseCase creates a new DeliveryUseCase.
func NewDeliveryUseCase(repo DeliveryRepository) *DeliveryUseCase {
	return &DeliveryUseCase{repo: repo}
}

// SaveDriverPosition saves a driver position.
func (uc *DeliveryUseCase) SaveDriverPosition(driverID uint64, latitude, longitude float64) error {
	return uc.repo.SaveDriverPosition(driverID, latitude, longitude)
}

// FindDriverPosition finds a driver position.
func (uc *DeliveryUseCase) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {
	return uc.repo.FindDriverPosition(driverID)
}

// HistoryDriverPosition finds a driver position history.
func (uc *DeliveryUseCase) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	return uc.repo.HistoryDriverPosition(driverID)
}
