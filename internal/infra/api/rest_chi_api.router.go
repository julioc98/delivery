package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/julioc98/delivery/internal/domain"
)

// DeliveryUseCase represents a use case for delivery drivers.
type DeliveryUseCase interface {
	// Save saves a driver position.
	SaveDriverPosition(driverID uint64, latitude, longitude float64) error
	// FindDriverPosition finds a driver position.
	FindDriverPosition(driverID uint64) (domain.DriverPosition, error)
	// HistoryDriverPosition finds a driver position history.
	HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error)
	// GetDriversNearby finds drivers nearby.
	GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error)
}

// DeliveryRestHandler represents a REST handler for delivery drivers.
type DeliveryRestHandler struct {
	r  *chi.Mux
	uc DeliveryUseCase
}

// NewDeliveryRestHandler creates a new DeliveryRestHandler.
func NewDeliveryRestHandler(r *chi.Mux, uc DeliveryUseCase) *DeliveryRestHandler {
	return &DeliveryRestHandler{
		r:  r,
		uc: uc,
	}
}

// RegisterHandlers registers the handlers of the REST API.
func (h *DeliveryRestHandler) RegisterHandlers() {
	h.r.Post("/drivers/{driverID}/locations", h.SavePosition)
	h.r.Get("/drivers/{driverID}/locations/current", h.FindPosition)
	h.r.Get("/drivers/{driverID}/locations", h.HistoryPosition)
	h.r.Get("/drivers/locations/nearby", h.FindNearby)
}
