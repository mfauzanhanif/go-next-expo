# Aplikasi — Multi-Tenant SaaS Platform

Monorepo untuk platform SaaS dengan arsitektur **Go Backend + Next.js Frontend + React Native Mobile**, diorkestrasi menggunakan Docker Compose.

---

## Daftar Isi

- [Prasyarat](#prasyarat)
- [Langkah 1 — Clone Repository](#langkah-1--clone-repository)
- [Langkah 2 — Setup Environment Variables](#langkah-2--setup-environment-variables)
- [Langkah 3 — Install Dependensi Frontend](#langkah-3--install-dependensi-frontend)
- [Langkah 4 — Install Dependensi Backend](#langkah-4--install-dependensi-backend)
- [Langkah 5 — Jalankan Aplikasi](#langkah-5--jalankan-aplikasi)
- [Akses Layanan](#akses-layanan)
- [Perintah Makefile](#perintah-makefile)
- [Struktur Proyek](#struktur-proyek)
- [Troubleshooting](#troubleshooting)

---

## Prasyarat

Pastikan semua tools berikut sudah terinstall di mesin Anda **sebelum** memulai:

| Tool | Versi Minimum | Cara Install | Keterangan |
| :--- | :--- | :--- | :--- |
| **Docker Desktop** | 4.x | [docker.com/get-docker](https://docs.docker.com/get-docker/) | Wajib. Pastikan Docker Engine berjalan. |
| **Docker Compose** | v2+ | Sudah termasuk di Docker Desktop | Plugin bawaan Docker Desktop. |
| **GNU Make** | 4.x | `sudo apt install make` | Untuk menjalankan perintah Makefile. |
| **Git** | 2.x | `sudo apt install git` | Version control. |
| **Go** | 1.26+ | [go.dev/dl](https://go.dev/dl/) | Untuk development backend lokal. |
| **Node.js** | v24+ | [nodejs.org](https://nodejs.org/) atau via `nvm` | Runtime JavaScript untuk frontend. |
| **PNPM** | 11.1.1 | `sudo npm install -g pnpm` | Package manager frontend (monorepo). |

### Verifikasi Instalasi

Jalankan perintah berikut untuk memastikan semua tools terinstall:

```bash
docker --version        # Docker version 28.x+
docker compose version  # Docker Compose version v2.x+
make --version          # GNU Make 4.x
git --version           # git version 2.x
go version              # go1.26.x
node --version          # v24.x
pnpm --version          # 11.1.1
```

---

## Langkah 1 — Clone Repository

```bash
git clone <URL_REPOSITORY> aplikasi
cd aplikasi
```

---

## Langkah 2 — Setup Environment Variables

File `.env` di-ignore oleh `.gitignore` untuk keamanan. Anda perlu membuatnya secara manual dari template yang disediakan.

### Backend

```bash
cp backend/.env.example backend/.env
```

Isi file `backend/.env`:

```env
# PostgreSQL Database
POSTGRES_USER=aplikasi
POSTGRES_PASSWORD=secret
POSTGRES_DB=aplikasi_db

# Redis
REDIS_PASSWORD=secret

# MinIO
MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin

# Backend App
BACKEND_PORT=8080
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8081
```

> **Penting:** Untuk environment **produksi**, ganti semua nilai default di atas dengan password yang kuat dan aman.

---

## Langkah 3 — Install Dependensi Frontend

Folder `node_modules/` dan cache Turborepo (`.turbo/`) di-ignore oleh `.gitignore`, sehingga harus di-generate ulang setelah clone.

```bash
cd frontend
pnpm install
cd ..
```

Perintah ini akan:
- Membaca `pnpm-lock.yaml` yang sudah ada di repository
- Mengunduh seluruh dependensi ke `node_modules/` di setiap workspace (`web`, `mobile`, `packages/*`)
- Men-setup Turborepo untuk build orchestration

---

## Langkah 4 — Install Dependensi Backend

Folder `vendor/` (jika ada) dan binary Go di-ignore. Jalankan:

```bash
cd backend
go mod download
cd ..
```

Atau jika Anda ingin menggunakan vendor:

```bash
cd backend
go mod vendor
cd ..
```

---

## Langkah 5 — Jalankan Aplikasi

### Opsi A: Jalankan semua via Docker (Rekomendasi)

```bash
make up
```

Perintah ini akan mem-build dan menjalankan **seluruh stack**: PostgreSQL, Redis, Zitadel, MinIO, Backend, Worker, Frontend, Caddy, Prometheus, Grafana, dan OpenTelemetry Collector.

### Opsi B: Development lokal (Frontend & Backend di luar Docker)

```bash
# Jalankan hanya infrastruktur via Docker
make dev

# Di terminal terpisah, jalankan backend
cd backend && make run

# Di terminal lain, jalankan frontend
cd frontend && make dev-web
```

---

## Akses Layanan

Setelah semua container berjalan, layanan dapat diakses di:

| Layanan | URL | Keterangan |
| :--- | :--- | :--- |
| **Frontend (Next.js)** | http://localhost:3000 | Aplikasi web utama |
| **Backend API (Go)** | http://localhost:8080 | REST API |
| **PostgreSQL** | `localhost:5432` | Database (koneksi via client) |
| **Redis** | `localhost:6379` | Cache & queue |
| **Zitadel** | http://localhost:8085 | Identity provider / IAM |
| **MinIO Console** | http://localhost:9001 | Object storage dashboard |
| **MinIO API** | `localhost:9000` | S3-compatible API |
| **Prometheus** | http://localhost:9090 | Monitoring & metrics |
| **Grafana** | http://localhost:3001 | Dashboard visualisasi (admin/admin) |
| **Caddy (Reverse Proxy)** | http://localhost:80 | Entry point produksi |

---

## Perintah Makefile

Berikut ringkasan perintah yang paling sering digunakan:

```bash
make help            # Tampilkan semua perintah yang tersedia
make up              # Jalankan semua layanan
make down            # Hentikan semua layanan
make restart         # Restart semua layanan
make ps              # Lihat status container
make logs            # Lihat logs semua layanan
make logs-backend    # Lihat logs backend & worker
make logs-frontend   # Lihat logs frontend
make logs-infra      # Lihat logs infrastruktur

make up-infra        # Jalankan infra saja (DB, Redis, Zitadel, MinIO)
make up-app          # Jalankan app saja (Backend, Worker, Frontend)
make up-monitoring   # Jalankan monitoring saja

make dev             # Mode development (infra Docker + app lokal)
make clean           # Hapus semua container & volumes
make status          # Lihat status lengkap
```

Untuk dokumentasi lengkap semua perintah Makefile (termasuk backend dan frontend), lihat [docs/makefile.md](docs/makefile.md).

---

## Struktur Proyek

```
aplikasi/
├── backend/               # Go + Echo v5 (API server & Worker)
│   ├── cmd/
│   │   ├── api/           # Entry point API server
│   │   └── worker/        # Entry point Asynq worker
│   ├── .env.example       # Template environment variables
│   ├── Dockerfile         # Multi-stage Docker build
│   ├── Makefile            # Backend-specific commands
│   ├── go.mod
│   └── go.sum
├── frontend/              # Turborepo monorepo (PNPM workspace)
│   ├── web/               # Next.js web app
│   ├── mobile/            # React Native + Expo (segera)
│   ├── packages/          # Shared packages (UI, utils, schemas)
│   ├── Dockerfile         # Multi-stage Docker build (Next.js)
│   ├── Makefile            # Frontend-specific commands
│   ├── turbo.json         # Turborepo pipeline config
│   ├── pnpm-workspace.yaml
│   ├── pnpm-lock.yaml
│   └── package.json
├── infra/                 # Konfigurasi infrastruktur
│   ├── caddy/
│   │   └── Caddyfile      # Reverse proxy config
│   └── monitoring/
│       ├── prometheus.yaml
│       ├── otel-collector.yaml
│       └── grafana/       # Grafana provisioning
├── docs/                  # Dokumentasi proyek
├── docker-compose.yaml    # Orkestrasi semua layanan
├── Makefile               # Root orchestration commands
├── .gitignore
└── README.md
```

---

## Troubleshooting

### "pnpm: command not found"

Node.js dan PNPM belum terinstall. Jalankan:

```bash
# Install Node.js v24
curl -fsSL https://deb.nodesource.com/setup_24.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install PNPM
sudo npm install -g pnpm
```

### "container aplikasi-postgres is unhealthy"

Volume PostgreSQL lama tidak kompatibel dengan versi baru. Hapus volume dan mulai ulang:

```bash
docker compose down
docker volume rm aplikasi_postgres_data
make up
```

### "port is already allocated"

Port sudah digunakan oleh proses lain. Cek proses yang menggunakan port:

```bash
sudo lsof -i :<PORT_NUMBER>
```

Atau ubah port mapping di `docker-compose.yaml`.

### "turbo prune — Cannot prune without parsed lockfile"

File `pnpm-lock.yaml` belum ada. Jalankan:

```bash
cd frontend
pnpm install
cd ..
```

### Setelah pull/clone, frontend tidak bisa build

Semua folder `node_modules/` dan `.next/` di-ignore oleh Git. Selalu jalankan `pnpm install` terlebih dahulu:

```bash
cd frontend && pnpm install && cd ..
```
