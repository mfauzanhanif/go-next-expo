package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"

	"backend/internal/config"
	"backend/internal/ent"
	"backend/pkg/logger"
	"backend/pkg/response"
	appvalidator "backend/pkg/validator"
)

func main() {
	// =========================================================================
	// 1. Load Configuration
	// =========================================================================
	cfg := config.Load()

	// =========================================================================
	// 2. Initialize Logger
	// =========================================================================
	log := logger.New(cfg.App.Env)
	slog.SetDefault(log)

	log.Info("starting application",
		slog.String("app", cfg.App.Name),
		slog.String("env", cfg.App.Env),
		slog.String("port", cfg.App.Port),
	)

	// =========================================================================
	// 3. Initialize Infrastructure
	// =========================================================================
	db, err := config.NewDatabase(cfg.Database, log)
	if err != nil {
		log.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	entClient := config.NewEntClient(db, log)
	// Pastikan close dipanggil saat aplikasi berhenti
	defer entClient.Close()

	rdb, err := config.NewRedis(cfg.Redis, log)
	if err != nil {
		log.Error("failed to connect to redis", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer rdb.Close()

	mc, err := config.NewMinIO(cfg.MinIO, log)
	if err != nil {
		log.Error("failed to connect to minio", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// =========================================================================
	// 4. Initialize Echo
	// =========================================================================
	e := echo.New()
	e.Logger = log
	e.Validator = appvalidator.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(cfg.App.AllowedOrigins, ","),
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowCredentials: true,
	}))

	// =========================================================================
	// 5. Register Routes
	// =========================================================================
	registerRoutes(e, entClient, db, rdb, mc)

	// =========================================================================
	// 6. Start Server
	// =========================================================================
	server := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      e,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("server started", slog.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server failed to start", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	// =========================================================================
	// 7. Graceful Shutdown
	// =========================================================================
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Info("shutdown signal received, draining connections...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("server stopped gracefully")
}

// registerRoutes mendaftarkan semua route API.
// Dependensi (db, redis, minio) di-inject di sini agar mudah di-test.
func registerRoutes(e *echo.Echo, entClient *ent.Client, db *sql.DB, rdb *redis.Client, mc *minio.Client) {
	// Root — Informasi dasar API
	e.GET("/", func(c *echo.Context) error {
		return response.Success(c, http.StatusOK, map[string]string{
			"name":    "Aplikasi API",
			"version": "0.1.0",
			"status":  "running",
		})
	})

	// Health Check — Digunakan oleh Docker, Kubernetes, dan load balancer
	e.GET("/health", func(c *echo.Context) error {
		ctx := c.Request().Context()
		health := map[string]string{
			"status": "ok",
		}

		// Cek koneksi database
		if err := db.PingContext(ctx); err != nil {
			health["database"] = "error"
			health["status"] = "degraded"
		} else {
			health["database"] = "ok"
		}

		// Cek koneksi Redis
		if err := rdb.Ping(ctx).Err(); err != nil {
			health["redis"] = "error"
			health["status"] = "degraded"
		} else {
			health["redis"] = "ok"
		}

		code := http.StatusOK
		if health["status"] == "degraded" {
			code = http.StatusServiceUnavailable
		}

		return response.Success(c, code, health)
	})
}
