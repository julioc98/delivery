// Package app implements the application layer.
package app

// DeliveryRepository represents a repository for delivery drivers.
type DeliveryRepository interface {
	// SaveDriverPosition saves a driver position.
	SaveDriverPosition(driverID uint64, latitude, longitude float64) error
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
