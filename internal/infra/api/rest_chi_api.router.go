package api

import "github.com/go-chi/chi/v5"

// DeliveryUseCase represents a use case for delivery drivers.
type DeliveryUseCase interface {
	// Save saves a driver position.
	SaveDriverPosition(driverID uint64, latitude, longitude float64) error
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
}
