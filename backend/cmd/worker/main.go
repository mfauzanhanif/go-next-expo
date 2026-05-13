package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/hibiken/asynq"
)

func main() {
	log.Println("🚀 Worker is starting...")

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	// Parse REDIS_URL to Extract Address and Password
	// Format expected: redis://:password@host:port/db
	redisHost := "localhost:6379"
	redisPassword := ""

	if strings.HasPrefix(redisURL, "redis://") {
		trimmed := strings.TrimPrefix(redisURL, "redis://")
		parts := strings.Split(trimmed, "@")
		if len(parts) == 2 {
			// Extract password (assuming format :password)
			credParts := strings.Split(parts[0], ":")
			if len(credParts) == 2 {
				redisPassword = credParts[1]
			}
			
			// Extract host
			hostParts := strings.Split(parts[1], "/")
			redisHost = hostParts[0]
		} else {
			// No password, just host
			hostParts := strings.Split(trimmed, "/")
			redisHost = hostParts[0]
		}
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisHost, Password: redisPassword},
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
	// TODO: Register your tasks here
	// mux.HandleFunc("email:send", handler.HandleEmailDeliveryTask)

	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run worker server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Worker is shutting down...")
	srv.Shutdown()
	log.Println("✅ Worker successfully shutdown.")
}
