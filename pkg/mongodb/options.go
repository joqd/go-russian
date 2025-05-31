package mongodb

import "time"

// Option
type Option func(*MongoDB)

// WithConnAttempts sets the number of connection attempts
func WithConnAttempts(n int) Option {
	return func(m *MongoDB) {
		m.connAttempts = n
	}
}

// WithConnTimeout sets the connection timeout
func WithConnTimeout(d time.Duration) Option {
	return func(m *MongoDB) {
		m.connTimeout = d
	}
}

// WithShutdownTimeout sets the shutdown timeout
func WithShutdownTimeout(d time.Duration) Option {
	return func(m *MongoDB) {
		m.shutdownTimeout = d
	}
}
