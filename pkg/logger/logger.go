package logger

import (
	"log/slog"
	"os"
)

// NewLogger creates a JSON logger that writes to stdout.
func NewLogger(config Config) (*slog.Logger, error) {
	level, err := config.validate()
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})

	return slog.New(handler).With(
		slog.String("service", config.Service),
		slog.String("environment", config.Environment),
	), nil
}
