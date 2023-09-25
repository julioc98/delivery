package db

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/domain"
	"github.com/redis/go-redis/v9"
)

const (
	cachePrefix = "delivery:"
	cacheTTL    = 1 * time.Hour // Cache expiration time (1 hour, adjust as needed)
)

// RedisRepositoryDecorator is a repository that uses Redis as a database.
type RedisRepositoryDecorator struct {
	redisClient *redis.Client
	dbRepo      app.DeliveryRepository
}

// NewRedisRepositoryDecorator creates a new RedisRepositoryDecorator.
func NewRedisRepositoryDecorator(redisClient *redis.Client, dbRepo app.DeliveryRepository) *RedisRepositoryDecorator {
	return &RedisRepositoryDecorator{redisClient: redisClient, dbRepo: dbRepo}
}

// SaveDriverPosition saves a driver position.
func (repo *RedisRepositoryDecorator) SaveDriverPosition(driverID uint64, latitude, longitude float64) error {
	err := repo.dbRepo.SaveDriverPosition(driverID, latitude, longitude)
	if err != nil {
		return err
	}

	_ = repo.ClearCache()

	return nil
}

// FindDriverPosition retrieves a driver's position from either the Redis cache or the database.
func (repo *RedisRepositoryDecorator) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {
	// Cache key for the driver's position
	cacheKey := cachePrefix + "position:" + strconv.FormatUint(driverID, 10)

	// Try to get the position from the Redis cache
	cacheValue, err := repo.redisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		// Cache hit
		var position domain.DriverPosition
		if err := json.Unmarshal([]byte(cacheValue), &position); err == nil {
			return position, nil
		}
	}

	// Cache miss, get the position from the database
	position, err := repo.dbRepo.FindDriverPosition(driverID)
	if err != nil {
		return domain.DriverPosition{}, err
	}

	// Save the position in the cache as a JSON string
	cacheValueBytes, err := json.Marshal(position)
	if err != nil {
		return domain.DriverPosition{}, err
	}

	cacheValue = string(cacheValueBytes) // Convert bytes to string

	err = repo.redisClient.Set(context.Background(), cacheKey, cacheValue, cacheTTL).Err()
	if err != nil {
		log.Printf("Error saving cache key %s: %s", cacheKey, err)
	}

	return position, nil
}

// HistoryDriverPosition finds a driver position history.
func (repo *RedisRepositoryDecorator) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	// Cache key for the driver's position history
	cacheKey := cachePrefix + "history:" + strconv.FormatUint(driverID, 10)

	// Try to get the position history from the Redis cache
	cacheValue, err := repo.redisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		// Cache hit
		var history []domain.DriverPosition
		if err := json.Unmarshal([]byte(cacheValue), &history); err == nil {
			return history, nil
		}
	}

	// Cache miss, get the position history from the database
	history, err := repo.dbRepo.HistoryDriverPosition(driverID)
	if err != nil {
		return nil, err
	}

	// Save the position history in the cache as a JSON string
	cacheValueBytes, err := json.Marshal(history)
	if err != nil {
		return nil, err
	}

	cacheValue = string(cacheValueBytes) // Convert bytes to string

	err = repo.redisClient.Set(context.Background(), cacheKey, cacheValue, cacheTTL).Err()
	if err != nil {
		log.Printf("Error saving cache key %s: %s", cacheKey, err)
	}

	return history, nil
}

// GetDriversNearby finds drivers nearby.
func (repo *RedisRepositoryDecorator) GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error) {
	// Cache key for nearby drivers based on latitude, longitude, and radius
	cacheKey := cachePrefix + "nearby:" +
		strconv.FormatFloat(latitude, 'f', -1, 64) + ":" +
		strconv.FormatFloat(longitude, 'f', -1, 64) + ":" +
		strconv.Itoa(radius)

	// Try to get nearby drivers from the Redis cache
	cacheValue, err := repo.redisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		// Cache hit
		var nearbyDrivers []domain.DriverPosition
		if err := json.Unmarshal([]byte(cacheValue), &nearbyDrivers); err == nil {
			return nearbyDrivers, nil
		}
	}

	// Cache miss, get nearby drivers from the database
	nearbyDrivers, err := repo.dbRepo.GetDriversNearby(latitude, longitude, radius)
	if err != nil {
		return nil, err
	}

	// Save nearby drivers in the cache as a JSON string
	cacheValueBytes, err := json.Marshal(nearbyDrivers)
	if err != nil {
		return nil, err
	}

	cacheValue = string(cacheValueBytes) // Convert bytes to string

	err = repo.redisClient.Set(context.Background(), cacheKey, cacheValue, cacheTTL).Err()
	if err != nil {
		log.Printf("Error saving cache key %s: %s", cacheKey, err)
	}

	return nearbyDrivers, nil
}

// ClearCache clears all caches.
func (repo *RedisRepositoryDecorator) ClearCache() error {
	// Implement logic to clear all cache keys
	keys, err := repo.redisClient.Keys(context.Background(), cachePrefix+"*").Result()
	if err != nil {
		return err
	}

	// Delete cache keys
	if len(keys) > 0 {
		_, err := repo.redisClient.Del(context.Background(), keys...).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
