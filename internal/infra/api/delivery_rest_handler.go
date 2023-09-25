// Package api implements the API layer.
package api

import (
	"database/sql"
	"encoding/json"
	"errors"
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

	if err := h.cmd.SaveDriverPosition(driverIDUint64, position.Lat, position.Long); err != nil {
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

	position, err := h.qry.FindDriverPosition(driverIDUint64)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Driver not found", http.StatusNotFound)

			return
		}

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

	positions, err := h.qry.HistoryDriverPosition(driverIDUint64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if len(positions) == 0 {
		http.Error(w, "Driver not found", http.StatusNotFound)

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

// FindNearby finds drivers nearby.
func (h *DeliveryRestHandler) FindNearby(w http.ResponseWriter, r *http.Request) {
	// Get the location coordinates and radius from the query parameters
	latitudeParam := r.URL.Query().Get("latitude")
	longitudeParam := r.URL.Query().Get("longitude")
	radiusParam := r.URL.Query().Get("radius")

	// Convert latitude, longitude, and radius to appropriate data types
	latitude, err := strconv.ParseFloat(latitudeParam, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)

		return
	}

	longitude, err := strconv.ParseFloat(longitudeParam, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)

		return
	}

	radius, err := strconv.Atoi(radiusParam)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)

		return
	}

	positions, err := h.qry.GetDriversNearby(latitude, longitude, radius)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if len(positions) == 0 {
		w.WriteHeader(http.StatusNotFound)

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
