package logger

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// Config for logger, use va
type Config struct {
	Level       string `envconfig:"LOG_LEVEL" default:"info"`
	Service     string `envconfig:"SERVICE_NAME" required:"true"`
	Environment string `envconfig:"APP_ENV" default:"local"`
}

func Load() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("loading config: %w", err)
	}
	if _, err := config.validate(); err != nil {
		return Config{}, fmt.Errorf("validate logger config: %w", err)
	}
	return config, nil
}

func MustLoad() Config {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	return config
}

// validate checks the configuration and returns the parsed log level.
func (c Config) validate() (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(c.Level)); err != nil {
		return 0, fmt.Errorf("parse log level: %w", err)
	}
	if strings.TrimSpace(c.Service) == "" {
		return 0, fmt.Errorf("service name must not be empty")
	}
	if strings.TrimSpace(c.Environment) == "" {
		return 0, fmt.Errorf("environment must not be empty")
	}
	return level, nil
}
