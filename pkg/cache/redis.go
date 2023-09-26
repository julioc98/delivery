// Package cache provides a Redis client.
package cache

import "github.com/redis/go-redis/v9"

// RedisConn creates a Redis client.
func RedisConn() (*redis.Client, error) {
	opt, err := redis.ParseURL("redis://@redis:6379/0")
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	return client, nil
}
