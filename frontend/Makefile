# =============================================================================
# Frontend Makefile — Turborepo + PNPM (Next.js Web & React Native Mobile)
# Command shortcut untuk development frontend
# =============================================================================

# Variabel
PNPM        := pnpm
TURBO       := pnpm turbo

# Warna output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RED    := \033[0;31m
NC     := \033[0m # No Color

.PHONY: help install dev dev-web dev-mobile build build-web lint lint-web lint-mobile \
        format type-check clean docker-build

# Default target
.DEFAULT_GOAL := help

# =============================================================================
# HELP
# =============================================================================

help: ## Tampilkan daftar command yang tersedia
	@echo "$(GREEN)Frontend Makefile Commands:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'
	@echo ""

# =============================================================================
# SETUP
# =============================================================================

install: ## Install semua dependensi workspace
	@echo "$(GREEN)>> Installing dependensi...$(NC)"
	$(PNPM) install

install-frozen: ## Install dependensi tanpa update lockfile
	@echo "$(GREEN)>> Installing dependensi (frozen lockfile)...$(NC)"
	$(PNPM) install --frozen-lockfile

# =============================================================================
# DEVELOPMENT
# =============================================================================

dev: ## Jalankan semua apps (Web + Mobile) secara bersamaan
	@echo "$(GREEN)>> Menjalankan semua apps...$(NC)"
	$(TURBO) dev

dev-web: ## Jalankan Next.js web (development)
	@echo "$(GREEN)>> Menjalankan Next.js web...$(NC)"
	$(TURBO) dev --filter=web

dev-mobile: ## Jalankan Expo mobile (development)
	@echo "$(GREEN)>> Menjalankan Expo mobile...$(NC)"
	$(TURBO) dev --filter=mobile

# =============================================================================
# BUILD
# =============================================================================

build: ## Build semua apps via Turborepo
	@echo "$(GREEN)>> Building semua apps...$(NC)"
	$(TURBO) build

build-web: ## Build Next.js web saja
	@echo "$(GREEN)>> Building Next.js web...$(NC)"
	$(TURBO) build --filter=web

# =============================================================================
# QUALITY ASSURANCE
# =============================================================================

lint: ## Jalankan linter untuk semua packages
	@echo "$(GREEN)>> Menjalankan linter...$(NC)"
	$(TURBO) lint

lint-web: ## Jalankan linter untuk web saja
	@echo "$(GREEN)>> Menjalankan linter web...$(NC)"
	$(TURBO) lint --filter=web

lint-mobile: ## Jalankan linter untuk mobile saja
	@echo "$(GREEN)>> Menjalankan linter mobile...$(NC)"
	$(TURBO) lint --filter=mobile

format: ## Format semua kode dengan Prettier
	@echo "$(GREEN)>> Formatting kode...$(NC)"
	$(PNPM) prettier --write .

format-check: ## Cek format kode tanpa mengubah
	@echo "$(GREEN)>> Checking format kode...$(NC)"
	$(PNPM) prettier --check .

type-check: ## Jalankan TypeScript type checking
	@echo "$(GREEN)>> Type checking...$(NC)"
	$(TURBO) type-check

test: ## Jalankan semua tests
	@echo "$(GREEN)>> Menjalankan tests...$(NC)"
	$(TURBO) test

test-web: ## Jalankan tests untuk web saja
	@echo "$(GREEN)>> Menjalankan tests web...$(NC)"
	$(TURBO) test --filter=web

check: lint type-check test ## Jalankan semua quality checks (lint, type-check, test)

# =============================================================================
# EXPO / MOBILE
# =============================================================================

expo-start: ## Jalankan Expo development server
	@echo "$(GREEN)>> Menjalankan Expo...$(NC)"
	cd mobile && $(PNPM) expo start

expo-ios: ## Jalankan Expo di iOS Simulator
	@echo "$(GREEN)>> Menjalankan Expo iOS...$(NC)"
	cd mobile && $(PNPM) expo start --ios

expo-android: ## Jalankan Expo di Android Emulator
	@echo "$(GREEN)>> Menjalankan Expo Android...$(NC)"
	cd mobile && $(PNPM) expo start --android

expo-prebuild: ## Generate native project files
	@echo "$(GREEN)>> Running Expo prebuild...$(NC)"
	cd mobile && $(PNPM) expo prebuild

# =============================================================================
# PACKAGES (Shared)
# =============================================================================

build-packages: ## Build shared packages
	@echo "$(GREEN)>> Building shared packages...$(NC)"
	$(TURBO) build --filter='./packages/*'

# =============================================================================
# DOCKER
# =============================================================================

docker-build: ## Build Docker image frontend (web)
	@echo "$(GREEN)>> Building Docker image frontend...$(NC)"
	docker build -t aplikasi-frontend:latest .

# =============================================================================
# CLEANUP
# =============================================================================

clean: ## Bersihkan semua build artifacts dan cache
	@echo "$(GREEN)>> Membersihkan...$(NC)"
	@rm -rf node_modules
	@rm -rf web/.next
	@rm -rf web/node_modules
	@rm -rf mobile/node_modules
	@rm -rf packages/*/node_modules
	@rm -rf .turbo
	@rm -rf web/.turbo
	@rm -rf mobile/.turbo
	@rm -rf mobile/.expo
	@echo "$(GREEN)>> Bersih!$(NC)"

clean-cache: ## Bersihkan Turborepo cache saja
	@echo "$(GREEN)>> Membersihkan Turbo cache...$(NC)"
	@rm -rf .turbo
	$(TURBO) clean
	@echo "$(GREEN)>> Cache bersih!$(NC)"