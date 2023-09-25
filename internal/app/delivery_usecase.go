// Package app implements the application layer.
package app

import (
	"fmt"
	"log"

	"github.com/julioc98/delivery/internal/domain"
)

// NewPositionSubject represents a new position subject.
const NewPositionSubject = "new.position"

// DeliveryUseCase represents a use case for delivery drivers.
type DeliveryUseCase struct {
	repo  DeliveryRepository
	event Publisher
}

// NewDeliveryUseCase creates a new DeliveryUseCase.
func NewDeliveryUseCase(repo DeliveryRepository, event Publisher) *DeliveryUseCase {
	return &DeliveryUseCase{repo: repo, event: event}
}

// SaveDriverPosition saves a driver position.
func (uc *DeliveryUseCase) SaveDriverPosition(driverID uint64, latitude, longitude float64) error {
	id, err := uc.repo.SaveDriverPosition(driverID, latitude, longitude)
	if err != nil {
		return err
	}

	err = uc.event.Publish(NewPositionSubject, []byte(fmt.Sprintf(`{"id":%d}`, id)))
	if err != nil {
		log.Println("publish err:", err, "id:", id)
	}

	return nil
}

// FindDriverPosition finds a driver position.
func (uc *DeliveryUseCase) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {
	return uc.repo.FindDriverPosition(driverID)
}

// HistoryDriverPosition finds a driver position history.
func (uc *DeliveryUseCase) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	return uc.repo.HistoryDriverPosition(driverID)
}

// GetDriversNearby finds drivers nearby.
func (uc *DeliveryUseCase) GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error) {
	return uc.repo.GetDriversNearby(latitude, longitude, radius)
}
