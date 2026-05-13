# =============================================================================
# Root Makefile — Orkestrasi Semua Layanan
# Command shortcut utama untuk mengelola seluruh stack aplikasi
# =============================================================================

# Variabel
COMPOSE         := docker compose
COMPOSE_FILE    := docker-compose.yaml
COMPOSE_PROD    := -f $(COMPOSE_FILE) -f infra/docker/docker-compose.prod.yaml

# Warna output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RED    := \033[0;31m
CYAN   := \033[0;36m
NC     := \033[0m # No Color

.PHONY: help up down restart ps logs build pull \
        up-infra up-app up-monitoring \
        backend-% frontend-% \
        dev dev-backend dev-frontend \
        clean prune deploy status

# Default target
.DEFAULT_GOAL := help

# =============================================================================
# HELP
# =============================================================================

help: ## Tampilkan daftar command yang tersedia
	@echo ""
	@echo "$(CYAN)╔══════════════════════════════════════════════════════╗$(NC)"
	@echo "$(CYAN)║          APLIKASI — Orkestrasi Makefile             ║$(NC)"
	@echo "$(CYAN)╚══════════════════════════════════════════════════════╝$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_%-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-25s$(NC) %s\n", $$1, $$2}'
	@echo ""

# =============================================================================
# DOCKER COMPOSE — Full Stack
# =============================================================================

up: ## Jalankan semua layanan (build & detached)
	@echo "$(GREEN)>> Menjalankan semua layanan...$(NC)"
	$(COMPOSE) up -d --build

down: ## Hentikan semua layanan
	@echo "$(GREEN)>> Menghentikan semua layanan...$(NC)"
	$(COMPOSE) down

restart: down up ## Restart semua layanan

ps: ## Tampilkan status semua container
	@$(COMPOSE) ps

logs: ## Tampilkan logs semua layanan (follow)
	$(COMPOSE) logs -f

logs-backend: ## Tampilkan logs backend (follow)
	$(COMPOSE) logs -f backend worker

logs-frontend: ## Tampilkan logs frontend (follow)
	$(COMPOSE) logs -f frontend

logs-infra: ## Tampilkan logs infrastruktur (follow)
	$(COMPOSE) logs -f postgres redis zitadel minio caddy

build: ## Build ulang semua Docker images
	@echo "$(GREEN)>> Building semua images...$(NC)"
	$(COMPOSE) build

pull: ## Pull versi terbaru semua base images
	@echo "$(GREEN)>> Pulling base images...$(NC)"
	$(COMPOSE) pull

# =============================================================================
# DOCKER COMPOSE — Selective Services
# =============================================================================

up-infra: ## Jalankan layanan infrastruktur saja (DB, Cache, Auth, Storage)
	@echo "$(GREEN)>> Menjalankan infrastruktur...$(NC)"
	$(COMPOSE) up -d postgres redis zitadel minio

up-app: ## Jalankan layanan aplikasi saja (Backend, Worker, Frontend)
	@echo "$(GREEN)>> Menjalankan aplikasi...$(NC)"
	$(COMPOSE) up -d backend worker frontend

up-monitoring: ## Jalankan layanan monitoring saja (Prometheus, Grafana, OTel)
	@echo "$(GREEN)>> Menjalankan monitoring stack...$(NC)"
	$(COMPOSE) up -d prometheus grafana otel-collector

up-backend: ## Jalankan backend + dependensi (DB, Cache, Auth)
	@echo "$(GREEN)>> Menjalankan backend + dependensi...$(NC)"
	$(COMPOSE) up -d postgres redis zitadel minio backend worker

up-frontend: ## Jalankan frontend + backend + dependensi
	@echo "$(GREEN)>> Menjalankan frontend + backend...$(NC)"
	$(COMPOSE) up -d postgres redis zitadel minio backend worker frontend

# =============================================================================
# DEVELOPMENT LOKAL (tanpa Docker untuk app)
# =============================================================================

dev: ## Jalankan infra (Docker) + backend & frontend (lokal)
	@echo "$(GREEN)>> Menjalankan infra via Docker...$(NC)"
	$(COMPOSE) up -d postgres redis zitadel minio
	@echo "$(GREEN)>> Jalankan backend dan frontend secara terpisah:$(NC)"
	@echo "  $(YELLOW)Terminal 1:$(NC) cd backend && make run"
	@echo "  $(YELLOW)Terminal 2:$(NC) cd frontend && make dev-web"

dev-backend: ## Jalankan infra + backend lokal
	@echo "$(GREEN)>> Menjalankan infra via Docker...$(NC)"
	$(COMPOSE) up -d postgres redis zitadel minio
	@echo "$(GREEN)>> Menjalankan backend secara lokal...$(NC)"
	$(MAKE) -C backend run

dev-frontend: ## Jalankan frontend lokal (pastikan backend sudah jalan)
	@echo "$(GREEN)>> Menjalankan frontend secara lokal...$(NC)"
	$(MAKE) -C frontend dev-web

# =============================================================================
# DELEGASI KE SUB-MAKEFILE
# =============================================================================

backend-%: ## Delegasi command ke backend Makefile (contoh: make backend-test)
	@$(MAKE) -C backend $*

frontend-%: ## Delegasi command ke frontend Makefile (contoh: make frontend-lint)
	@$(MAKE) -C frontend $*

# =============================================================================
# DATABASE
# =============================================================================

db-migrate: ## Jalankan migrasi database
	@$(MAKE) -C backend migrate

db-status: ## Cek status migrasi database
	@$(MAKE) -C backend migrate-status

# =============================================================================
# PRODUCTION
# =============================================================================

deploy: ## Deploy ke produksi (via docker-compose.prod.yaml)
	@echo "$(GREEN)>> Deploying ke produksi...$(NC)"
	$(COMPOSE) $(COMPOSE_PROD) up -d --build
	@echo "$(GREEN)>> Deploy selesai!$(NC)"

deploy-down: ## Hentikan deployment produksi
	@echo "$(GREEN)>> Menghentikan deployment produksi...$(NC)"
	$(COMPOSE) $(COMPOSE_PROD) down

# =============================================================================
# STATUS & HEALTH
# =============================================================================

status: ## Tampilkan status lengkap semua layanan
	@echo ""
	@echo "$(CYAN)=== Container Status ===$(NC)"
	@$(COMPOSE) ps
	@echo ""
	@echo "$(CYAN)=== Disk Usage ===$(NC)"
	@docker system df 2>/dev/null || true
	@echo ""
	@echo "$(CYAN)=== Volume Usage ===$(NC)"
	@docker volume ls --filter name=aplikasi 2>/dev/null || true

# =============================================================================
# CLEANUP
# =============================================================================

clean: ## Hentikan layanan dan hapus volumes
	@echo "$(RED)>> Menghentikan dan membersihkan semua data...$(NC)"
	$(COMPOSE) down -v --remove-orphans
	@echo "$(GREEN)>> Bersih!$(NC)"

prune: ## Bersihkan Docker resources yang tidak terpakai
	@echo "$(RED)>> Membersihkan Docker resources...$(NC)"
	docker system prune -f
	docker volume prune -f
	@echo "$(GREEN)>> Docker bersih!$(NC)"
