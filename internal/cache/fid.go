// Cache for storing detected flights and for performing
// optimization on flight recall.
package cache

import (
	"context"
	"errors"
	"time"

	"github.com/kingcobra2468/atracker/internal/radar"

	"github.com/go-redis/redis/v8"
)

// Manage a Redis client instance and context.
type fidRedisCache struct {
	rdb *redis.Client
	ctx context.Context
}

var fidCache *fidRedisCache

// Establish a new connection instance with Redis.
func Connect(addr string) {
	if fidCache != nil {
		return
	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	context := client.Context()

	fidCache = &fidRedisCache{rdb: client, ctx: context}
}

// Check if a given fid is already being tracked. If true,
// this means that the fid is already recorded (until it expires).
func CheckPlaneTracked(fid radar.FlightID) (bool, error) {
	if fidCache == nil {
		return false, errors.New("redis connection not established")
	}

	switch err := fidCache.rdb.Get(fidCache.ctx, fid.ID).Err(); err {
	case redis.Nil:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

// Cache a given fid and set the cache to expire in 5 minutes. This
// will also reset the TTL if called on an fid that already exists.
func TrackNew(fid radar.FlightID) error {
	if fidCache == nil {
		return errors.New("redis connection not established")
	}
	dur := time.Duration(time.Minute * 5)

	return fidCache.rdb.Set(fidCache.ctx, fid.ID, fid.ID, dur).Err()
}

// Only cache a fid if it is not currently being cached.
func TrackIfNew(fid radar.FlightID) error {
	if fidCache == nil {
		return errors.New("redis connection not established")
	}

	new, err := CheckPlaneTracked(fid)
	if err != nil {
		return err
	}
	if !new {
		return nil
	}

	return TrackNew(fid)
}
