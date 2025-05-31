// Package mongodb implements MongoDB connection.
package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	_defaultConnAttempts    = 10
	_defaultConnTimeout     = 1 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// MongoDB
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database

	connAttempts    int
	connTimeout     time.Duration
	shutdownTimeout time.Duration
}

// New
func New(uri, dbName string, opts ...Option) (*MongoDB, error) {
	m := &MongoDB{
		connAttempts:    _defaultConnAttempts,
		connTimeout:     _defaultConnTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(m)
	}

	clientOpts := options.Client().ApplyURI(uri)

	var err error
	for m.connAttempts > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), m.connTimeout)
		defer cancel()

		m.Client, err = mongo.Connect(clientOpts)
		if err == nil {
			err = m.Client.Ping(ctx, nil)
			if err == nil {
				break
			}
		}

		log.Printf("MongoDB is trying to connect, attempts left: %d", m.connAttempts-1)
		m.connAttempts--
		time.Sleep(m.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("mongodb - New - connection failed: %w", err)
	}

	m.Database = m.Client.Database(dbName)

	return m, nil
}

// Close
func (m *MongoDB) Close(ctx context.Context) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, m.shutdownTimeout)
	defer cancel()

	if m.Client != nil {
		return m.Client.Disconnect(ctxWithTimeout)
	}
	return nil
}
