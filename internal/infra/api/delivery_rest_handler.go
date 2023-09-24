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
