package redisdb

import "time"

// Option
type Option func(*Redis)

// WithConnAttempts sets connection retry attempts
func WithConnAttempts(n int) Option {
	return func(r *Redis) {
		r.connAttempts = n
	}
}

// WithConnTimeout sets connection timeout
func WithConnTimeout(d time.Duration) Option {
	return func(r *Redis) {
		r.connTimeout = d
	}
}

// WithConnTimeout sets connection timeout
func WithTTL(d time.Duration) Option {
	return func(r *Redis) {
		r.TTL = d
	}
}
