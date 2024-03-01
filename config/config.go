package config

import "time"

type (
	Config struct {
		HTTP
		PostgreSQL
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
	}
)
