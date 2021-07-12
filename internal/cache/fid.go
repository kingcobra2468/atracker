// Cache for storing detected flights and for performing
// optimization on flight recall.
package cache

import (
	"context"
	"time"

	"github.com/kingcobra2468/atracker/internal/radar"

	"github.com/go-redis/redis/v8"
)

// Manage a Redis client instance and context.
type FidRedisCache struct {
	Addr string
	rdb  *redis.Client
	ctx  context.Context
}

// Establish a new connection instance with Redis.
func (frc *FidRedisCache) Connect() {
	frc.rdb = redis.NewClient(&redis.Options{
		Addr:     frc.Addr,
		Password: "",
		DB:       0,
	})
	frc.ctx = frc.rdb.Context()
}

// Check if a given fid is already being tracked. If true,
// this means that the fid is already recorded (until it expires).
func (frc *FidRedisCache) CheckPlaneTracked(fid radar.FlightID) (bool, error) {
	switch err := frc.rdb.Get(frc.ctx, fid.ID).Err(); err {
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
func (frc *FidRedisCache) TrackNew(fid radar.FlightID) error {
	dur := time.Duration(time.Minute * 5)

	return frc.rdb.Set(frc.ctx, fid.ID, fid.ID, dur).Err()
}

// Only cache and fid if it is not currently being cached.
func (frc *FidRedisCache) TrackIfNew(fid radar.FlightID) error {
	new, err := frc.CheckPlaneTracked(fid)
	if err != nil {
		return err
	}
	if !new {
		return nil
	}

	return frc.TrackNew(fid)
}
