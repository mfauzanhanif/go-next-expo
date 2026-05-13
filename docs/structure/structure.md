# STRUCTURE

/aplikasi
├── .github # CI/CD Pipeline (GitHub Actions)
│ └── /workflows  
│ ├── build-and-test.yaml # CI: Menjalankan linter, test, dan build (Frontend & Backend)
│ └── deploy-production.yaml # CD: Otomatisasi rilis ke server produksi menggunakan Docker
├── /backend # Repositori Backend (Go + Fiber v3)
│ ├── Makefile # Command shortcut - Spesialis Go - Migrasi, generate, build API
│ └── Dockerfile # Dockerfile untuk backend
├── /frontend # Repositori Frontend (Turborepo Workspace)
│ ├── /web # Repositori Frontend Web (Next.js)
│ ├── /mobile # Repositori Frontend Mobile (React Native)
│ ├── /packages # Folder berbagi UI/Zod nantinya
│ ├── Makefile # Command shortcut - Spesialis UI - Menjalankan Next.js, Expo, Turborepo
│ └── Dockerfile # Dockerfile untuk frontend
├── /infra # Konfigurasi Infrastruktur
│ ├── /caddy # Caddyfile & on-demand TLS script
│ ├── /docker # Docker Compose & Swarm configs
│ ├── /minio # MinIO configs
│ └── /monitoring # Prometheus & Grafana configs
├── docker-compose.yaml  
└── Makefile # Command shortcut - Orkestrasi semua layanan
