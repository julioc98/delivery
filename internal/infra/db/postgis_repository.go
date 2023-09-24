// Package db implements the database layer.
package db

import (
	"database/sql"
	"fmt"
)

// DeliveryPostGISRepository represents a repository for delivery drivers.
type DeliveryPostGISRepository struct {
	db *sql.DB
}

// NewDeliveryPostGISRepository creates a new DeliveryPostGISRepository.
func NewDeliveryPostGISRepository(db *sql.DB) *DeliveryPostGISRepository {
	return &DeliveryPostGISRepository{db: db}
}

// SaveDriverPosition saves a driver position.
func (repo *DeliveryPostGISRepository) SaveDriverPosition(driverID uint64, latitude, longitude float64) error {
	query := "INSERT INTO driver_positions (driver_id, location) VALUES ($1, ST_GeomFromText($2, 4326))"
	coordinates := fmt.Sprintf("POINT(%f %f)", longitude, latitude)
	_, err := repo.db.Exec(query, driverID, coordinates)

	return err
}
