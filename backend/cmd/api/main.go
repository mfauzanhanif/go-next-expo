package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	// Echo v5 mengganti Logger() dengan Logger-basis Slog atau middleware logging baru
	// Kita akan menggunakan middleware standar
	e.Use(middleware.Recover())

	// Strict CORS Configuration
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000,http://localhost:8081" // Default fallback dev
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(allowedOrigins, ","),
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	// Routes
	// Di Echo v5, echo.Context adalah *echo.Context (pointer)
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"message": "Super App Backend is running smoothly on Echo v5",
		})
	})

	// Pengaturan Port
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	// Buat standar http.Server untuk graceful shutdown yang lebih robust di Echo v5
	server := &http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	// 1. Jalankan server di dalam Goroutine agar tidak memblokir eksekusi di bawahnya
	go func() {
		log.Printf("🚀 Server API berjalan di port :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Gagal memulai server: %v", err)
		}
	}()

	// 2. Graceful Shutdown Logic (Menangkap sinyal dari OS/Docker)
	quit := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C), SIGTERM (Sinyal terminasi dari Docker/Kubernetes)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	<-quit // Menunggu sampai sinyal masuk
	log.Println("🛑 Menerima sinyal shutdown, sedang mematikan server secara perlahan...")

	// Memberikan batas waktu 10 detik agar request yang sedang diproses bisa selesai dulu
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Gagal mematikan server secara perlahan: %v", err)
	}

	log.Println("✅ Server berhasil dimatikan.")
}