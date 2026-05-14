package logger

import (
	"log/slog"
	"os"
)

// New membuat instance slog.Logger baru yang dikonfigurasi berdasarkan environment.
// - "development": Text handler dengan level Debug (mudah dibaca di terminal)
// - "production" dan lainnya: JSON handler dengan level Info (optimal untuk log aggregator)
func New(env string) *slog.Logger {
	var handler slog.Handler

	switch env {
	case "development":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	return slog.New(handler)
}
