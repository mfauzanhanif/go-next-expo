package config

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"backend/internal/database"
	"backend/internal/ent"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDatabase membuat koneksi PostgreSQL menggunakan driver pgx
// dengan connection pool yang dikonfigurasi sesuai best practice produksi.
func NewDatabase(cfg DatabaseConfig, logger *slog.Logger) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("config: gagal membuka koneksi database: %w", err)
	}

	// Konfigurasi connection pool
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetimeMin) * time.Minute)

	// Verifikasi koneksi dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("config: gagal ping database: %w", err)
	}

	logger.Info("database connected",
		slog.String("driver", "pgx"),
		slog.Int("max_open_conns", cfg.MaxOpenConns),
		slog.Int("max_idle_conns", cfg.MaxIdleConns),
	)

	return db, nil
}

// NewEntClient membuat instance ent.Client baru yang sudah di-wrap dengan tenantDriver
// untuk mendukung dynamic schema switching (multi-tenancy).
func NewEntClient(db *sql.DB, logger *slog.Logger) *ent.Client {
	drv := database.NewTenantDriver(db)

	opts := []ent.Option{
		ent.Driver(drv),
	}

	// Di mode development, bisa ditambahkan logger untuk query
	// opts = append(opts, ent.Debug())

	return ent.NewClient(opts...)
}
