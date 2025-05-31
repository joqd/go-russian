// Package redisdb implements Redis connection.
package redisdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = 1 * time.Second
	_defaultTTL          = 1 * time.Hour
)

// Redis
type Redis struct {
	Client *redis.Client
	TTL    time.Duration

	connAttempts int
	connTimeout  time.Duration
}

// New
func New(url string, opts ...Option) (*Redis, error) {
	r := &Redis{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		TTL:          _defaultTTL,
	}

	for _, opt := range opts {
		opt(r)
	}

	var err error
	urlOpts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	r.Client = redis.NewClient(urlOpts)

	for r.connAttempts > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), r.connTimeout)

		_, err = r.Client.Ping(ctx).Result()
		cancel()

		if err == nil {
			break
		}

		log.Printf("Redis is trying to connect, attempts left: %d", r.connAttempts-1)
		r.connAttempts--
		time.Sleep(r.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("redisdb - New - connection failed: %w", err)
	}

	return r, nil
}

// Close -.
func (r *Redis) Close() error {
	if r.Client != nil {
		return r.Client.Close()
	}
	return nil
}
