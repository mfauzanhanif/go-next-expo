# =============================================================================
# Backend Makefile — Go + Echo v5
# Command shortcut untuk development backend lokal
# =============================================================================

# Gunakan bash sebagai shell (diperlukan untuk `source`, `set -a`, dll.)
SHELL := /bin/bash

# Pastikan PATH menyertakan lokasi go (snap) dan atlas
export PATH := /snap/bin:/usr/local/bin:$(PATH)

# Variabel
APP_NAME        := api-server
WORKER_NAME     := worker
GO              := go
GOFLAGS         := -v
MAIN_API        := ./cmd/api
MAIN_WORKER     := ./cmd/worker
BUILD_DIR       := ./bin
ENV_FILE        := .env

# Warna output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RED    := \033[0;31m
NC     := \033[0m # No Color

.PHONY: help run run-worker build build-worker test test-coverage lint fmt vet \
        tidy vendor swagger migrate migrate-status migrate-new \
        ent-generate ent-new clean docker-build

# Default target
.DEFAULT_GOAL := help

# =============================================================================
# HELP
# =============================================================================

help: ## Tampilkan daftar command yang tersedia
	@echo -e "$(GREEN)Backend Makefile Commands:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'
	@echo ""

# =============================================================================
# DEVELOPMENT
# =============================================================================

run: ## Jalankan API server (development)
	@echo -e "$(GREEN)>> Menjalankan API server...$(NC)"
	$(GO) run $(GOFLAGS) $(MAIN_API)

run-worker: ## Jalankan Asynq worker (development)
	@echo -e "$(GREEN)>> Menjalankan Asynq worker...$(NC)"
	$(GO) run $(GOFLAGS) $(MAIN_WORKER)

run-all: ## Jalankan API + Worker secara bersamaan
	@echo -e "$(GREEN)>> Menjalankan API server dan Worker...$(NC)"
	@trap 'kill 0' EXIT; \
		$(GO) run $(GOFLAGS) $(MAIN_API) & \
		$(GO) run $(GOFLAGS) $(MAIN_WORKER) & \
		wait

# =============================================================================
# BUILD
# =============================================================================

build: ## Build binary API server
	@echo -e "$(GREEN)>> Building API server...$(NC)"
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_API)
	@echo -e "$(GREEN)>> Binary: $(BUILD_DIR)/$(APP_NAME)$(NC)"

build-worker: ## Build binary Worker
	@echo -e "$(GREEN)>> Building Worker...$(NC)"
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -o $(BUILD_DIR)/$(WORKER_NAME) $(MAIN_WORKER)
	@echo -e "$(GREEN)>> Binary: $(BUILD_DIR)/$(WORKER_NAME)$(NC)"

build-all: build build-worker ## Build semua binary (API + Worker)

# =============================================================================
# QUALITY ASSURANCE
# =============================================================================

test: ## Jalankan semua unit tests
	@echo -e "$(GREEN)>> Menjalankan tests...$(NC)"
	$(GO) test ./... -race -count=1

test-coverage: ## Jalankan tests dengan coverage report
	@echo -e "$(GREEN)>> Menjalankan tests dengan coverage...$(NC)"
	@mkdir -p $(BUILD_DIR)
	$(GO) test ./... -race -count=1 -coverprofile=$(BUILD_DIR)/coverage.out
	$(GO) tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html
	@echo -e "$(GREEN)>> Coverage report: $(BUILD_DIR)/coverage.html$(NC)"

lint: ## Jalankan golangci-lint
	@echo -e "$(GREEN)>> Menjalankan linter...$(NC)"
	golangci-lint run ./...

fmt: ## Format semua kode Go
	@echo -e "$(GREEN)>> Formatting kode...$(NC)"
	$(GO) fmt ./...
	goimports -w .

vet: ## Jalankan go vet untuk analisis statis
	@echo -e "$(GREEN)>> Menjalankan go vet...$(NC)"
	$(GO) vet ./...

check: fmt vet lint test ## Jalankan semua quality checks (format, vet, lint, test)

# =============================================================================
# DEPENDENCY MANAGEMENT
# =============================================================================

tidy: ## Bersihkan dan sinkronisasi dependensi Go
	@echo -e "$(GREEN)>> Menjalankan go mod tidy...$(NC)"
	$(GO) mod tidy

vendor: ## Download dependensi ke folder vendor
	@echo -e "$(GREEN)>> Vendoring dependensi...$(NC)"
	$(GO) mod vendor

# =============================================================================
# CODE GENERATION
# =============================================================================

swagger: ## Generate dokumentasi Swagger/OpenAPI via Swaggo
	@echo -e "$(GREEN)>> Generating Swagger docs...$(NC)"
	~/go/bin/swag init -g cmd/api/main.go -o ./docs --parseDependency --parseInternal
	@echo -e "$(GREEN)>> Swagger docs generated di ./docs$(NC)"

ent-generate: ## Generate kode Ent ORM dari skema
	@echo -e "$(GREEN)>> Generating Ent code...$(NC)"
	$(GO) generate ./internal/ent
	@echo -e "$(GREEN)>> Ent code generated$(NC)"

ent-new: ## Buat skema Ent baru (usage: make ent-new NAME=User)
	@if [ -z "$(NAME)" ]; then \
		echo "$(RED)>> Error: Tentukan nama skema. Contoh: make ent-new NAME=User$(NC)"; \
		exit 1; \
	fi
	@echo -e "$(GREEN)>> Membuat skema Ent baru: $(NAME)...$(NC)"
	$(GO) run -mod=mod entgo.io/ent/cmd/ent new $(NAME)

# =============================================================================
# DATABASE MIGRATION (Atlas)
# =============================================================================

migrate: ## Jalankan migrasi database
	@echo -e "$(GREEN)>> Menjalankan migrasi database...$(NC)"
	@set -a; . $(ENV_FILE); set +a; atlas migrate apply --env local --url "$$DATABASE_URL"

migrate-status: ## Cek status migrasi database
	@echo -e "$(GREEN)>> Cek status migrasi...$(NC)"
	@set -a; . $(ENV_FILE); set +a; atlas migrate status --env local --url "$$DATABASE_URL"

migrate-new: ## Buat file migrasi baru (usage: make migrate-new NAME=add_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "$(RED)>> Error: Tentukan nama migrasi. Contoh: make migrate-new NAME=add_users_table$(NC)"; \
		exit 1; \
	fi
	@echo -e "$(GREEN)>> Membuat migrasi baru: $(NAME)...$(NC)"
	@set -a; . $(ENV_FILE); set +a; atlas migrate diff $(NAME) --env local

migrate-hash: ## Rehash migrasi (setelah edit manual)
	@echo -e "$(GREEN)>> Rehashing migrasi...$(NC)"
	atlas migrate hash --env local

db-tables: ## Lihat daftar tabel di database (via psql)
	@echo -e "$(GREEN)>> Menampilkan daftar tabel...$(NC)"
	docker exec -i aplikasi-postgres psql -U aplikasi -d aplikasi_db -c '\dt'

db-desc: ## Lihat struktur tabel (usage: make db-desc TABLE=nama_tabel)
	@if [ -z "$(TABLE)" ]; then \
		echo "$(RED)>> Error: Tentukan nama tabel. Contoh: make db-desc TABLE=tenants$(NC)"; \
		exit 1; \
	fi
	@echo -e "$(GREEN)>> Menampilkan struktur tabel $(TABLE)...$(NC)"
	docker exec -i aplikasi-postgres psql -U aplikasi -d aplikasi_db -c '\d $(TABLE)'

migration-history: ## Lihat riwayat migrasi di database (via psql)
	@echo -e "$(GREEN)>> Menampilkan riwayat migrasi...$(NC)"
	docker exec -i aplikasi-postgres psql -U aplikasi -d aplikasi_db -c 'SELECT * FROM atlas_schema_revisions.atlas_schema_revisions;'

# =============================================================================
# DOCKER
# =============================================================================

docker-build: ## Build Docker image backend
	@echo -e "$(GREEN)>> Building Docker image backend...$(NC)"
	docker build -t aplikasi-backend:latest .

docker-run: ## Jalankan container backend
	@echo -e "$(GREEN)>> Menjalankan container backend...$(NC)"
	docker run --rm -p 8080:8080 --env-file $(ENV_FILE) aplikasi-backend:latest

# =============================================================================
# CLEANUP
# =============================================================================

clean: ## Bersihkan file build dan cache
	@echo -e "$(GREEN)>> Membersihkan...$(NC)"
	@rm -rf $(BUILD_DIR)
	$(GO) clean -cache -testcache
	@echo -e "$(GREEN)>> Bersih!$(NC)"
