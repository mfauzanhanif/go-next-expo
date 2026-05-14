package config

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewMinIO membuat koneksi MinIO client dan memastikan default bucket tersedia.
func NewMinIO(cfg MinIOConfig, logger *slog.Logger) (*minio.Client, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("config: gagal membuat minio client: %w", err)
	}

	// Pastikan default bucket tersedia
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("config: gagal cek bucket '%s': %w", cfg.Bucket, err)
	}

	if !exists {
		if err := client.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("config: gagal membuat bucket '%s': %w", cfg.Bucket, err)
		}
		logger.Info("minio bucket created", slog.String("bucket", cfg.Bucket))
	}

	logger.Info("minio connected",
		slog.String("endpoint", cfg.Endpoint),
		slog.String("bucket", cfg.Bucket),
	)

	return client, nil
}
