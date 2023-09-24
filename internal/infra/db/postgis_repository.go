// Package db implements the database layer.
package db

import (
	"database/sql"
	"fmt"

	"github.com/julioc98/delivery/internal/domain"
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

// FindDriverPosition finds a driver position.
func (repo *DeliveryPostGISRepository) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {
	query := `SELECT id, driver_id, ST_X(location::geometry), ST_Y(location::geometry), timestamp 
		FROM driver_positions WHERE driver_id = $1 ORDER BY timestamp DESC LIMIT 1`
	row := repo.db.QueryRow(query, driverID)

	var dp domain.DriverPosition
	if err := row.Scan(&dp.ID, &dp.DriverID, &dp.Location.Long, &dp.Location.Lat, &dp.Timestamp); err != nil {
		return domain.DriverPosition{}, err
	}

	return dp, nil
}

// HistoryDriverPosition finds a driver position history.
func (repo *DeliveryPostGISRepository) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	query := `SELECT id, driver_id, ST_X(location::geometry), ST_Y(location::geometry), timestamp 
		FROM driver_positions WHERE driver_id = $1 ORDER BY timestamp DESC`

	rows, err := repo.db.Query(query, driverID)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var dps []domain.DriverPosition

	for rows.Next() {
		var dp domain.DriverPosition
		if err := rows.Scan(&dp.ID, &dp.DriverID, &dp.Location.Long, &dp.Location.Lat, &dp.Timestamp); err != nil {
			return nil, err
		}

		dps = append(dps, dp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dps, nil
}
