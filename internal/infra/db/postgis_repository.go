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

// GetDriversNearby finds drivers nearby.
func (repo *DeliveryPostGISRepository) GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error) {
	query := `
		WITH latest_driver_positions AS (
			SELECT driver_id, MAX(timestamp) AS max_timestamp
			FROM driver_positions
			GROUP BY driver_id
		)
		
		SELECT dp.id, dp.driver_id, ST_X(dp.location::geometry), ST_Y(dp.location::geometry), ldp.max_timestamp as timestamp
		FROM driver_positions dp
		JOIN latest_driver_positions ldp
		ON dp.driver_id = ldp.driver_id AND dp.timestamp = ldp.max_timestamp
		WHERE ST_DWithin(
			dp.location::geography,
			ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
			$3
		);
	`

	rows, err := repo.db.Query(query, longitude, latitude, radius)
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
