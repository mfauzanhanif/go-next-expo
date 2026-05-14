package config

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// NewRedis membuat koneksi Redis client dari URL format standar.
// Format URL: redis://:password@host:port/db
func NewRedis(cfg RedisConfig, logger *slog.Logger) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("config: gagal parse redis URL: %w", err)
	}

	client := redis.NewClient(opt)

	// Verifikasi koneksi dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("config: gagal ping redis: %w", err)
	}

	logger.Info("redis connected",
		slog.String("addr", opt.Addr),
		slog.Int("db", opt.DB),
	)

	return client, nil
}
