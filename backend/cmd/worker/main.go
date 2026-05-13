package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Worker is starting...")

	// TODO: Inisialisasi Asynq worker/server di sini
	// Contoh:
	// srv := asynq.NewServer(
	//     asynq.RedisClientOpt{Addr: os.Getenv("REDIS_URL")},
	//     asynq.Config{Concurrency: 10},
	// )

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Worker shutting down...")
}
