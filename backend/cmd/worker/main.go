package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"

	"backend/internal/config"
	"backend/pkg/logger"
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

	log.Info("starting worker",
		slog.String("app", cfg.App.Name),
		slog.String("env", cfg.App.Env),
	)

	// =========================================================================
	// 3. Parse Redis Connection
	// =========================================================================
	opt, err := redis.ParseURL(cfg.Redis.URL)
	if err != nil {
		log.Error("failed to parse redis URL", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// =========================================================================
	// 4. Initialize Asynq Server
	// =========================================================================
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     opt.Addr,
			Password: opt.Password,
			DB:       opt.DB,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	// TODO: Register task handlers di Fase 5
	// mux.HandleFunc("billing:generate", jobs.HandleBillingGenerate)
	// mux.HandleFunc("notification:send", jobs.HandleNotificationSend)

	// =========================================================================
	// 5. Start Worker
	// =========================================================================
	go func() {
		log.Info("worker started", slog.Int("concurrency", 10))
		if err := srv.Run(mux); err != nil {
			log.Error("worker failed to start", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	// =========================================================================
	// 6. Graceful Shutdown
	// =========================================================================
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("shutdown signal received, stopping worker...")
	srv.Shutdown()
	log.Info("worker stopped gracefully")
}
