// Package api implements the API layer.
package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/julioc98/delivery/internal/domain"
)

// SavePosition saves a driver position.
func (h *DeliveryRestHandler) SavePosition(w http.ResponseWriter, r *http.Request) {
	driverID := chi.URLParam(r, "driverID")

	driverIDUint64, err := strconv.ParseUint(driverID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid driverID", http.StatusBadRequest)

		return
	}

	var position domain.Point

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&position); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := h.uc.SaveDriverPosition(driverIDUint64, position.Lat, position.Long); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

// FindPosition finds a driver position.
func (h *DeliveryRestHandler) FindPosition(w http.ResponseWriter, r *http.Request) {
	driverID := chi.URLParam(r, "driverID")

	driverIDUint64, err := strconv.ParseUint(driverID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid driverID", http.StatusBadRequest)

		return
	}

	position, err := h.uc.FindDriverPosition(driverIDUint64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(position); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// HistoryPosition finds a driver position history.
func (h *DeliveryRestHandler) HistoryPosition(w http.ResponseWriter, r *http.Request) {
	driverID := chi.URLParam(r, "driverID")

	driverIDUint64, err := strconv.ParseUint(driverID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid driverID", http.StatusBadRequest)

		return
	}

	positions, err := h.uc.HistoryDriverPosition(driverIDUint64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(positions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
