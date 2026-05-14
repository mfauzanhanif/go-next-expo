package config

import (
	"os"
	"strconv"
)

// Config menyimpan seluruh konfigurasi aplikasi.
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	MinIO    MinIOConfig
}

// AppConfig berisi konfigurasi umum aplikasi.
type AppConfig struct {
	Name           string
	Env            string
	Port           string
	AllowedOrigins string
}

// DatabaseConfig berisi konfigurasi koneksi PostgreSQL.
type DatabaseConfig struct {
	URL                string
	MaxOpenConns       int
	MaxIdleConns       int
	ConnMaxLifetimeMin int
}

// RedisConfig berisi konfigurasi koneksi Redis.
type RedisConfig struct {
	URL string
}

// MinIOConfig berisi konfigurasi koneksi MinIO object storage.
type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
}

// Load membaca seluruh konfigurasi dari environment variables dengan fallback default.
func Load() *Config {
	return &Config{
		App: AppConfig{
			Name:           getEnv("APP_NAME", "aplikasi"),
			Env:            getEnv("APP_ENV", "development"),
			Port:           getEnv("BACKEND_PORT", "8080"),
			AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:8081"),
		},
		Database: DatabaseConfig{
			URL:                getEnv("DATABASE_URL", "postgres://aplikasi:secret@localhost:5432/aplikasi_db?sslmode=disable"),
			MaxOpenConns:       getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:       getEnvInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetimeMin: getEnvInt("DB_CONN_MAX_LIFETIME_MIN", 5),
		},
		Redis: RedisConfig{
			URL: getEnv("REDIS_URL", "redis://:secret@localhost:6379/0"),
		},
		MinIO: MinIOConfig{
			Endpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
			UseSSL:    getEnvBool("MINIO_USE_SSL", false),
			Bucket:    getEnv("MINIO_BUCKET", "aplikasi"),
		},
	}
}

// IsDevelopment mengembalikan true jika aplikasi berjalan di mode development.
func (c *Config) IsDevelopment() bool {
	return c.App.Env == "development"
}

// IsProduction mengembalikan true jika aplikasi berjalan di mode production.
func (c *Config) IsProduction() bool {
	return c.App.Env == "production"
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	return fallback
}
