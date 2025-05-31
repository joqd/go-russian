package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		Mongo   Mongo
		Redis   Redis
		Metrics Metrics
		Swagger Swagger
	}

	// App
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	// Log
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// Mongo
	Mongo struct {
		URI    string `env:"MONGO_URI,required"`
		DbName string `env:"MONGO_DB_NAME,required"`
	}

	// Redis
	Redis struct {
		URI string `env:"REDIS_URI,required"`
		TTL int    `env:"REDIS_TTL,required"`
	}

	// Metrics
	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	// Swagger
	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return conf, nil
}
