# Engineering Roadmap: Tahapan Pengembangan Backend

Dokumen ini menjabarkan rencana teknis pengembangan backend secara bertahap, dari fondasi arsitektur hingga integrasi worker dan penyimpanan objek. Setiap fase dipetakan langsung ke struktur direktori dan tech stack yang telah ditetapkan.

> Referensi: [PRD](prd.md) · [Tech Stack](techstack.md) · [Struktur Backend](structure/structure-backend.md) · [Skema Database](database/schema.dbml)

---

## Fase 1: Fondasi Arsitektur & Infrastruktur ✅

Fokus pada penyiapan kerangka kerja dasar, standardisasi kode, dan pengaturan lingkungan pengembangan.

### 1.1. Inisialisasi Proyek

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Go module (`go.mod`) | ✅ Selesai | `backend/go.mod` |
| Echo v5 HTTP Server | ✅ Selesai | `backend/cmd/api/main.go` |
| Graceful Shutdown (SIGINT/SIGTERM) | ✅ Selesai | `backend/cmd/api/main.go` |
| CORS Middleware | ✅ Selesai | `backend/cmd/api/main.go` |
| Asynq Worker Entry Point | ✅ Selesai | `backend/cmd/worker/main.go` |
| Health Check Route (`/health`) | ✅ Selesai | `backend/cmd/api/main.go` |

### 1.2. Standardisasi Utilitas (`/pkg`)

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Response JSON wrapper | ✅ Selesai | `backend/pkg/response/response.go` |
| Konfigurasi log terstruktur (Slog) | ✅ Selesai | `backend/pkg/logger/logger.go` |
| Integrasi Go-Playground Validator | ✅ Selesai | `backend/pkg/validator/validator.go` |

### 1.3. Koneksi Infrastruktur

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Loader konfigurasi (env vars) | ✅ Selesai | `backend/internal/config/config.go` |
| Connection pool PostgreSQL (pgx) | ✅ Selesai | `backend/internal/config/database.go` |
| Connection pool Redis | ✅ Selesai | `backend/internal/config/redis.go` |
| Koneksi MinIO SDK | ✅ Selesai | `backend/internal/config/minio.go` |

### 1.4. Kontainerisasi Lokal

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| PostgreSQL 18 (Alpine) | ✅ Selesai | `docker-compose.yaml` |
| Redis 8 (Alpine) | ✅ Selesai | `docker-compose.yaml` |
| MinIO (Object Storage) | ✅ Selesai | `docker-compose.yaml` |
| Caddy (Reverse Proxy) | ✅ Selesai | `docker-compose.yaml` |
| Prometheus + Grafana + OTel | ✅ Selesai | `docker-compose.yaml` |
| Zitadel (Identity Provider) | ✅ Selesai | `docker-compose.yaml` |
| Multi-stage Dockerfile (API + Worker) | ✅ Selesai | `backend/Dockerfile` |

---

## Fase 2: Multi-Tenancy & Skema Prioritas ✅

Membangun batas tegas antar lembaga menggunakan pendekatan *schema-per-tenant* dan mendefinisikan master data.

### 2.1. Skema Global — Ent ORM

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Entitas `Tenant` (public schema) | ✅ Selesai | `backend/internal/ent/schema/` |
| Entitas `User` (public schema) | ✅ Selesai | `backend/internal/ent/schema/` |
| Entitas `TenantUser` (relasi) | ✅ Selesai | `backend/internal/ent/schema/` |
| `go generate` untuk Ent | ✅ Selesai | `backend/internal/ent/generate.go` |

> **Referensi Skema:** Lihat tabel `public.tenants`, `public.users`, dan `public.tenant_users` di [schema.dbml](database/schema.dbml).

### 2.2. Skema Tenant — Ent ORM

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Entitas `GTK` (tenant schema) | ✅ Selesai | `backend/internal/ent/schema/` |
| Entitas `Department` (tenant schema) | ✅ Selesai | `backend/internal/ent/schema/` |
| Entitas `GTKPosition` (tenant schema) | ✅ Selesai | `backend/internal/ent/schema/` |

> **Referensi Skema:** Lihat tabel `tenant_xxx.gtk`, `tenant_xxx.departments`, dan `tenant_xxx.gtk_positions` di [schema.dbml](database/schema.dbml).

### 2.3. Migrasi Atlas

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Konfigurasi `atlas.hcl` | ✅ Selesai | `backend/atlas.hcl` |
| Skrip migrasi skema `public` | ✅ Selesai | `backend/internal/migrations/` |
| Automasi duplikasi skema per tenant baru | ✅ Selesai | `backend/internal/migrations/` |
| Integrasi target `make db-migrate` | ✅ Selesai | `backend/Makefile` |

### 2.4. Middleware Tenant

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Ekstraksi Tenant ID dari header/domain | ✅ Selesai | `backend/internal/middlewares/` |
| Injeksi Tenant ID ke context request | ✅ Selesai | `backend/internal/middlewares/` |
| Dynamic schema switching di Ent client | ✅ Selesai | `backend/internal/database/` |

---

## Fase 3: Autentikasi & Otorisasi Global (IAM)

Mengamankan endpoint API dan memastikan isolasi akses antar pengguna dan lembaga.

### 3.1. Integrasi Zitadel (AuthN)

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| OIDC token validation service | ⬜ Belum | `backend/internal/services/` |
| JWKS key fetching & caching | ⬜ Belum | `backend/internal/services/` |
| Middleware autentikasi Echo v5 | ⬜ Belum | `backend/internal/middlewares/` |

### 3.2. Engine Casbin (AuthZ)

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Definisi model Casbin (RBAC) | ⬜ Belum | `backend/internal/config/` |
| Adapter PostgreSQL untuk policy storage | ⬜ Belum | `backend/internal/config/` |
| Middleware otorisasi Echo v5 | ⬜ Belum | `backend/internal/middlewares/` |
| Sinkronisasi policy dengan lisensi modul | ⬜ Belum | `backend/internal/services/` |

### 3.3. Proteksi Endpoint

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Public routes (health, login callback) | ⬜ Belum | `backend/cmd/api/main.go` |
| Protected route groups (Zitadel + Casbin) | ⬜ Belum | `backend/cmd/api/main.go` |
| Rate limiting pada endpoint krusial | ⬜ Belum | `backend/internal/middlewares/` |

---

## Fase 4: Logika Bisnis & Eksekusi API (Master GTK)

Mengimplementasikan alur data dari database hingga menjadi response API yang siap dikonsumsi frontend Next.js.

### 4.1. Layer Data (`/internal/repositories`)

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Repository `Tenant` (CRUD) | ⬜ Belum | `backend/internal/repositories/` |
| Repository `GTK` (CRUD) | ⬜ Belum | `backend/internal/repositories/` |
| Repository `Department` (CRUD) | ⬜ Belum | `backend/internal/repositories/` |
| Repository `GTKPosition` (CRUD) | ⬜ Belum | `backend/internal/repositories/` |

### 4.2. Layer Bisnis (`/internal/services`)

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Service `Tenant` (provisioning skema baru) | ⬜ Belum | `backend/internal/services/` |
| Service `GTK` (validasi NIK/NUPTK unik) | ⬜ Belum | `backend/internal/services/` |
| Logika penugasan rangkap jabatan GTK | ⬜ Belum | `backend/internal/services/` |
| Multi-tenant context switching | ⬜ Belum | `backend/internal/services/` |

### 4.3. Layer Delivery (`/internal/handlers`)

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Handler `Tenant` (REST endpoints) | ⬜ Belum | `backend/internal/handlers/` |
| Handler `GTK` (REST endpoints) | ⬜ Belum | `backend/internal/handlers/` |
| Payload validation (Go-Playground) | ⬜ Belum | `backend/internal/handlers/` |
| Standardisasi response format | ⬜ Belum | `backend/internal/handlers/` |

### 4.4. Dokumentasi OpenAPI

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Anotasi Swaggo pada handlers | ⬜ Belum | `backend/internal/handlers/` |
| Generate Swagger JSON/YAML | ⬜ Belum | `backend/docs/` |
| Serve Swagger UI via Echo | ⬜ Belum | `backend/cmd/api/main.go` |
| Integrasi target `make swagger` | ⬜ Belum | `backend/Makefile` |

---

## Fase 5: Integrasi Worker & Penyimpanan Objek

Menangani tugas-tugas berat di luar request HTTP utama.

### 5.1. Setup MinIO

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| MinIO client wrapper | ✅ Selesai | `backend/internal/config/minio.go` |
| Service upload dokumen legal GTK | ⬜ Belum | `backend/internal/services/` |
| Service upload foto profil GTK | ⬜ Belum | `backend/internal/services/` |
| Handler API upload file | ⬜ Belum | `backend/internal/handlers/` |

### 5.2. Konfigurasi Asynq

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Entry point worker | ✅ Selesai | `backend/cmd/worker/main.go` |
| Task type definitions | ⬜ Belum | `backend/internal/jobs/` |
| Task handler implementations | ⬜ Belum | `backend/internal/jobs/` |
| Task enqueue helper | ⬜ Belum | `backend/internal/jobs/` |

### 5.3. Job Pertama

| Item | Status | Lokasi |
| :--- | :--- | :--- |
| Job: Sinkronisasi data massal | ⬜ Belum | `backend/internal/jobs/` |
| Job: Pengiriman notifikasi | ⬜ Belum | `backend/internal/jobs/` |
| Integrasi enqueue dari service layer | ⬜ Belum | `backend/internal/services/` |

---

## Ringkasan Status

| Fase | Deskripsi | Progress |
| :--- | :--- | :--- |
| **Fase 1** | Fondasi Arsitektur & Infrastruktur | ✅ Selesai |
| **Fase 2** | Multi-Tenancy & Skema Prioritas | 🟡 Sedang Dikerjakan |
| **Fase 3** | Autentikasi & Otorisasi Global | ⬜ Belum Dimulai |
| **Fase 4** | Logika Bisnis & Eksekusi API | ⬜ Belum Dimulai |
| **Fase 5** | Integrasi Worker & Penyimpanan Objek | 🟡 Entry Point + MinIO Selesai |
