# FRONTEND STRUCTURE

Struktur root dari direktori monorepo frontend (`/frontend`). Dokumentasi ini memberikan gambaran umum *top-level* file dan folder, serta detail khusus untuk direktori `packages`.

/frontend
├── /mobile                         # Direktori aplikasi React Native (Expo)
├── /web                            # Direktori aplikasi web Next.js (detail ada di structure-frontend-web.md)
├── /packages                       # Direktori shared workspace packages (digunakan lintas aplikasi)
│   ├── /config                     # Paket untuk konfigurasi bersama (misal: ESLint, TypeScript, Tailwind)
│   └── /ui                         # Paket untuk komponen antarmuka bersama (Shared UI / Shadcn)
├── Dockerfile                      # Konfigurasi multi-stage Docker build untuk frontend
├── Makefile                        # Shortcut command line untuk mempermudah development
├── package.json                    # Konfigurasi root package manager (Turborepo & PNPM)
├── pnpm-lock.yaml                  # Lockfile dependency yang men-track semua versi package
├── pnpm-workspace.yaml             # Definisi struktur PNPM workspace
└── turbo.json                      # Konfigurasi pipeline build system Turborepo
