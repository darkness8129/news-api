package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP
		PostgreSQL
		Test
	}

	HTTP struct {
		Addr            string        `env:"HTTP_ADDR" env-default:":8080"`
		WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT" env-default:"5s"`
		ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT" env-default:"5s"`
		ShutdownTimeout time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT" env-default:"3s"`
	}

	PostgreSQL struct {
		User     string `env:"POSTGRESQL_USER" env-default:"postgres"`
		Password string `env:"POSTGRESQL_PASSWORD" env-default:"postgres"`
		Host     string `env:"POSTGRESQL_HOST" env-default:"localhost"`
		Database string `env:"POSTGRESQL_DATABASE" env-default:"news_api"`
		Port     string `env:"POSTGRESQL_PORT" env-default:"5432"`
	}

	Test struct {
		PostgreSQLUser     string `env:"TEST_POSTGRESQL_USER" env-default:"postgres"`
		PostgreSQLPassword string `env:"TEST_POSTGRESQL_PASSWORD" env-default:"postgres"`
		PostgreSQLHost     string `env:"TEST_POSTGRESQL_HOST" env-default:"localhost"`
		PostgreSQLDatabase string `env:"TEST_POSTGRESQL_DATABASE" env-default:"test_news_api"`
		PostgreSQLPort     string `env:"TEST_POSTGRESQL_PORT" env-default:"5434"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return &cfg, nil
}
